package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

//nolint:funlen
func getNumericStats(
	ctx context.Context,
	client *ent.Client,
	gameVersion *ent.GameVersion,
) ([]*model.NumericMetric, error) {
	resp := make([]*model.NumericMetric, 0)
	rows, err := client.QueryContext(
		ctx,
		`WITH filtered_stats AS (
			SELECT
				sd.*,
				COALESCE(AVG(CAST(st.value AS NUMERIC)),
				0) AS average_value,
				COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (
			ORDER BY
				CAST(st.value AS NUMERIC)),
				0) AS median_value
			FROM
				stat_descriptions sd
			JOIN
				STATISTICS st ON
				sd.id = st.stat_description_stats
			JOIN
				stat_description_game_version sdgv ON
				sd.id = sdgv.stat_description_id
			JOIN
				game_versions gv ON
				gv.id = sdgv.game_version_id
			WHERE
				st.value ~ '^[+-]?[0-9]*\.?[0-9]+$'
				AND gv.id = $1
				AND (sd.type = 'numeric'
					OR sd.TYPE = 'aggregate')
			GROUP BY
				sd.id
			)
			SELECT
				fs.*
			FROM
				filtered_stats FS
			ORDER BY
				fs.order_number ASC;
		`,
		gameVersion.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var res struct {
			ent.StatDescription
			Metadata     *string
			AverageValue float64
			MedianValue  float64
		}
		err = rows.Scan(
			&res.ID,
			&res.Type,
			&res.Name,
			&res.Description,
			&res.Metadata,
			&res.OrderNumber,
			&res.AverageValue,
			&res.MedianValue,
		)
		if err != nil {
			return nil, err
		}
		metadata := ""
		if res.Metadata != nil {
			metadata = *res.Metadata
		}
		resp = append(resp, &model.NumericMetric{
			Stat: &ent.StatDescription{
				ID:          res.ID,
				Name:        res.Name,
				Description: res.Description,
				Type:        res.Type,
				OrderNumber: res.OrderNumber,
				Metadata:    metadata,
			},
			GlobalAverage: res.AverageValue,
		})
	}

	return resp, nil
}

//nolint:funlen
func getEnumStats(
	ctx context.Context,
	client *ent.Client,
	gameVersion *ent.GameVersion,
) ([]*model.EnumMetric, error) {
	resp := make([]*model.EnumMetric, 0)
	rows, err := client.QueryContext(
		ctx,
		`WITH filtered_stat_descriptions AS (
			SELECT
				sd.*
			FROM
				stat_descriptions sd
			JOIN
					stat_description_game_version sdgv ON
				sd.id = sdgv.stat_description_id
			JOIN
					game_versions gv ON
				gv.id = sdgv.game_version_id
			WHERE
				gv.id = $1
				AND sd."type" = 'enum'
			),
			count_aggregates AS (
			SELECT
				s.stat_description_stats AS stat_id,
				s.value,
				COUNT(*) AS occurrences
			FROM
				STATISTICS s
			JOIN
					filtered_stat_descriptions fsd ON
				s.stat_description_stats = fsd.id
			GROUP BY
				s.stat_description_stats,
				s.value
			),
			histogram_aggregates AS (
			SELECT
				stat_id,
				jsonb_object_agg(value,
				occurrences) AS histogram
			FROM
				count_aggregates
			GROUP BY
				stat_id
			)
			SELECT
				fsd.*,
				ha.histogram
			FROM
				filtered_stat_descriptions fsd
			LEFT JOIN
				histogram_aggregates ha ON
				fsd.id = ha.stat_id
			ORDER BY
				fsd.order_number ASC;`,
		gameVersion.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var res struct {
			ent.StatDescription
			Metadata  *string
			Histogram []uint8
		}
		err = rows.Scan(
			&res.ID,
			&res.Type,
			&res.Name,
			&res.Description,
			&res.Metadata,
			&res.OrderNumber,
			&res.Histogram,
		)
		if err != nil {
			return nil, err
		}
		metadata := ""
		if res.Metadata != nil {
			metadata = *res.Metadata
		}
		histogram := make(map[string]int)
		if len(res.Histogram) != 0 {
			err = json.Unmarshal(res.Histogram, &histogram)
			if err != nil {
				return nil, err
			}
		}
		occurrences := make([]*model.EnumOccurences, 0, len(histogram))
		for value, count := range histogram {
			occurrences = append(occurrences, &model.EnumOccurences{
				Value: value,
				//nolint:misspell
				Occurences: count,
			})
		}
		resp = append(resp, &model.EnumMetric{
			Stat: &ent.StatDescription{
				ID:          res.ID,
				Name:        res.Name,
				Description: res.Description,
				Type:        res.Type,
				OrderNumber: res.OrderNumber,
				Metadata:    metadata,
			},
			Global: occurrences,
		})
	}

	return resp, nil
}

var granularityToDuration = map[model.Granularity]time.Duration{
	//nolint:gomnd,mnd
	model.GranularityDay:   24 * time.Hour,
	model.GranularityWeek:  7 * 24 * time.Hour,
	model.GranularityMonth: 30 * 24 * time.Hour,
}

func getMatchesCreated(
	ctx context.Context,
	client *ent.Client,
	gameVersion *ent.GameVersion,
	input *model.TimeSeriesInput,
) (*model.TimeSeries, error) {
	period := strings.ToLower(input.Granularity.Value.String())
	resp := make([]*model.TimeSeriesPeriod, 0)
	rows, err := client.QueryContext(
		ctx,
		fmt.Sprintf(
			`WITH periods AS (SELECT generate_series($1::date, $2::date, '1 %s'::interval) AS period)
			SELECT
				p.period,
				COALESCE(COUNT(m.created_at), 0) AS occurrences
			FROM
				periods p
			LEFT JOIN
				matches m ON m.created_at >= p.period AND m.created_at < p.period + INTERVAL '1 %s' AND m.game_version_matches = $3
			GROUP BY
				p.period
			ORDER BY
				p.period;`,
			period, period),
		input.Start,
		input.End,
		gameVersion.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var res struct {
			Period      string
			Occurrences int
		}
		err = rows.Scan(&res.Period, &res.Occurrences)
		if err != nil {
			return nil, err
		}
		t, err := time.Parse(time.RFC3339, res.Period)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &model.TimeSeriesPeriod{
			ActivityCount: res.Occurrences,
			Start:         t,
			End:           t.Add(granularityToDuration[input.Granularity.Value]),
		})
	}

	return &model.TimeSeries{
		Series: resp,
	}, nil
}

func getAdoption(
	ctx context.Context,
	client *ent.Client,
	gameVersion *ent.GameVersion,
	input *model.GranularityInput,
) (*model.TimeFloatMetric, error) {
	return &model.TimeFloatMetric{
		Value: 0,
		Trend: 0,
	}, nil
}
