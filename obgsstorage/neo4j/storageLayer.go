package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/open-boardgame-stats/backend/obgsstorage"
)

// Neo4j ...
type Neo4j struct {
	session neo4j.Session
}

var _ obgsstorage.Storage = (*Neo4j)(nil)

// New ...
func New(target string) (*Neo4j, error) {
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }

	// driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "1234", ""), configForNeo4j40)
	driver, err := neo4j.NewDriver(target, neo4j.BasicAuth("neo4j", "1234", ""), configForNeo4j40)
	if err != nil {
		return nil, err
	}

	defer driver.Close()

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	return &Neo4j{session: session}, nil
}

// UserAdd ...
func (s *Neo4j) UserAdd(name, birth, gender string) error {
	result, err := s.session.Run("CREATE (:User  {name: $name, birth: date($birth), gender: $gender})", map[string]interface{}{
		"name":   name,
		"birth":  birth,
		"gender": gender,
	})

	if err != nil {
		return err
	}

	return result.Err()
}

// UserRemove ...
func (s *Neo4j) UserRemove(name string) error {

	result, err := s.session.Run("MATCH (u:User {name: $name})	DELETE u", map[string]interface{}{
		"name": name,
	})

	if err != nil {
		return err
	}

	return result.Err()
}

// UserRequestFriendship ...
func (s *Neo4j) UserRequestFriendship(nameFrom, nameTo string) error {
	// todo
	return nil
}

// UserRemoveFriendship ...
func (s *Neo4j) UserRemoveFriendship(nameFrom, nameTo string) error {
	// todo
	return nil
}

func testNeo4j() error {
	// configForNeo4j35 := func(conf *neo4j.Config) {}
	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }

	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "1234", ""), configForNeo4j40)
	if err != nil {
		return err
	}
	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	// For multidatabase support, set sessionConfig.DatabaseName to requested database
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		return err
	}
	defer session.Close()

	result, err := session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	if err != nil {
		return err
	}

	for result.Next() {
		fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	}
	return result.Err()
}
