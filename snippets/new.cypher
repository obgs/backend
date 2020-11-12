// Create constraints
CREATE CONSTRAINT ConstraintUserNode
ON (un:User)
ASSERT (un.name)
IS NODE KEY

CREATE CONSTRAINT ConstraintUserExistBirth
ON (ueb:User)
ASSERT EXISTS (ueb.birth)

CREATE CONSTRAINT ConstraintUserExistGender
ON (ueg:User)
ASSERT EXISTS (ueg.gender)

CREATE CONSTRAINT ConstraintGame
ON (g:Game)
ASSERT (g.name)
IS NODE KEY

// Create users
CREATE (u1:User  {name: 'Matt',     birth: date('1995-9-16'),  gender: 'male'})
CREATE (u2:User  {name: 'John',     birth: date('1998-3-14'),  gender: 'male'})
CREATE (u3:User  {name: 'Elijah',   birth: date('1979-11-22'), gender: 'male'})
CREATE (u4:User  {name: 'Mike',     birth: date('1996-5-6'),   gender: 'two-spirit'})
CREATE (u5:User  {name: 'Jeniffer', birth: date('2003-8-28'),  gender: 'agender'})
CREATE (u6:User  {name: 'Mary',     birth: date('2003-10-8'),  gender: 'female'})
CREATE (u7:User  {name: 'Adolf',    birth: date('1989-04-20'), gender: 'male'})
CREATE (u8:User  {name: 'Alex',     birth: date('2014-3-7'),   gender: 'attack helicopter'})
CREATE (u9:User  {name: 'Lyn',      birth: date('2013-10-27'), gender: 'female'})
CREATE (u10:User {name: 'Sam',      birth: date('1997-8-20'),  gender: 'female'})
CREATE (u1) -[:FRIEND]-> (u4)
CREATE (u7) -[:FRIEND]-> (u1)
CREATE (u7) -[:FRIEND]-> (u2)
CREATE (u7) -[:FRIEND]-> (u3)
CREATE (u4) -[:FRIEND]-> (u5)
CREATE (u8) -[:FRIEND]-> (u5)
CREATE (u9) -[:FRIEND]-> (u10)

// Friend requests
CALL {
  MATCH (requesting_user:User {name: 'Sam'})
  MATCH (requested_user:User {name: 'Mike'})
  CREATE (requesting_user) -[new_request:REQUESTED_FRIEND]-> (requested_user)
  RETURN requesting_user, requested_user, new_request
}
MATCH (requesting_user) <-[old_request:REQUESTED_FRIEND]- (requested_user)
DELETE new_request, old_request
CREATE (requesting_user) -[f:FRIEND]-> (requested_user)
RETURN id(new_request), id(f)

// User remove
MATCH (u:User {name: 'Mary'})
DELETE u
