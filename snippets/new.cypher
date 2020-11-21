// Create constraints
CREATE CONSTRAINT ConstraintUserNode
ON (un:User)
ASSERT (un.username)
IS NODE KEY

CREATE CONSTRAINT ConstraintGame
ON (g:Game)
ASSERT (g.name)
IS NODE KEY

CREATE CONSTRAINT ConstraintGameInstance
ON (m:GameInstance)
ASSERT (m.uuid)
IS NODE KEY

// Create users
CREATE (u1:User  {username: 'matt',     displayname: 'Matt',     birthdate: date('1995-9-16'),  gender: 'male'})
CREATE (u2:User  {username: 'john',     displayname: 'John',     birthdate: date('1998-3-14'),  gender: 'male'})
CREATE (u3:User  {username: 'elijah',   displayname: 'Elijah',   birthdate: date('1979-11-22'), gender: 'male'})
CREATE (u4:User  {username: 'mike',     displayname: 'Mike',     birthdate: date('1996-5-6'),   gender: 'two-spirit'})
CREATE (u5:User  {username: 'jeniffer', displayname: 'Jeniffer', birthdate: date('2003-8-28'),  gender: 'agender'})
CREATE (u6:User  {username: 'mary',     displayname: 'Mary',     birthdate: date('2003-10-8'),  gender: 'female'})
CREATE (u7:User  {username: 'adolf',    displayname: 'Adolf',    birthdate: date('1989-04-20'), gender: 'male'})
CREATE (u8:User  {username: 'alex',     displayname: 'Alex',     birthdate: date('2014-3-7'),   gender: 'attack helicopter'})
CREATE (u9:User  {username: 'lyn',      displayname: 'Lyn',      birthdate: date('2013-10-27'), gender: 'female'})
CREATE (u10:User {username: 'sam',      displayname: 'Sam',      birthdate: date('1997-8-20'),  gender: 'female'})
CREATE (u1) -[:FRIEND]-> (u4)
CREATE (u7) -[:FRIEND]-> (u1)
CREATE (u7) -[:FRIEND]-> (u2)
CREATE (u7) -[:FRIEND]-> (u3)
CREATE (u4) -[:FRIEND]-> (u5)
CREATE (u8) -[:FRIEND]-> (u5)
CREATE (u9) -[:FRIEND]-> (u10)

// Friend requests
CALL {
  MATCH (requesting_user:User {username: 'sam'})
  MATCH (requested_user:User {username: 'mike'})
  CREATE (requesting_user) -[new_request:REQUESTED_FRIEND]-> (requested_user)
  RETURN requesting_user, requested_user, new_request
}
MATCH (requesting_user) <-[old_request:REQUESTED_FRIEND]- (requested_user)
DELETE new_request, old_request
CREATE (requesting_user) -[f:FRIEND]-> (requested_user)
RETURN id(new_request), id(f)

// User remove
MATCH (u:User {name: 'mary'})
DELETE u

// Unfriend
MATCH (requesting_user:User {username: 'sam'}) -[r:FRIEND]- (requested_user:User {username: 'mike'})
CREATE (requesting_user) <-[:REQUESTED_FRIEND]- (requested_user)
DELETE r;
MATCH (requesting_user:User {username: 'sam'}) -[r:REQUESTED_FRIEND]-> (requested_user:User {username: 'mike'})
DELETE r

// Games
CREATE (:Game {
  name: 'terraforming-mars',
  metrics: [
    'terraform_rating',
    'achievements',
    'awards',
    'cards',
    'cities',
    'forests'
  ]
})

// GameInstance
MATCH (g:Game name: {'terraforming-mars'})
CREATE (m:GameInstance date: date('2020-11-16'))
  -[:OF_GAME {
    terraform_rating: 1, // неверно
    achievements:     2,
    awards:           3,
    cards:            4,
    cities:           5,
    forests:          6
  }]->
  (g)
