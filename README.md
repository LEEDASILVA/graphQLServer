# graphQL

documentation for [graphQL](https://spec.graphql.org/)

[graph-Q-L](https://graphql.org):

- new API standard that was invented & open-sourced by facebook its an alternative to REST
- enables declarative data fetching were a client can specify the right data it wants
- insted of multiple endpoints that return fixed data-structures, a graphQL server only exposes a single endpoint and respondes with the specific data the client asks ford

Why is graphQL the better REST?
for the REST we have to crete for example:
server side:

```console
    urls: [ /users/<id> /users/<id>/posts /users/<id>/followers ]

    all the urls get be fetched by the client

    GET request: the users name we have to fetch the link `</users/<id>`
        This will then load all the information about the user but we just needed the name -.-
    GET request: the users posts we have to fetch the link `/users/<id>/posts`
        Same thing here we load all the posts.... a lot of information
    ....
```

so we get the idea and the problems with the RESTfull API
problem:

- Overfetching

- Underfetching

But using graphQL:

we just have a single endpoint to be used by all the clients, so basically no specific paths

we will just need to send a single request to the server :D so just one request to get all (one to rule them all)

```console
    **POST** request:
    query {
        User(id: "34j3n1ej") {
            name
            post {
                title
            }
            followers(last: 3) {
                name
            }
        }
    }

```

## **core concepts**

schema definition Language(SDL)

- Defining simple types

```js
// ! -> required
type Person {
    id: ID!
    name: String!
    age: Int!
}

type Post {
    title: String!
}

// relation between the person and post one person can have several post
// one to many type
type Person {
    name: String!
    age: Int!
    posts: [Post!]!
}

type Post {
    title: String!
    author: Person!
}
```

example query that the client can send to the server:

- simple:

```js
// root filed of the query
{
    allPersons {
        name
    }
}
```

response

```json
{
    "allPersons": [
        { "name": "Johnny" },
        { "name": "Sarah" },
        { "name": "Alice" }
    ]
}
```

- last:

```js
// tag last is present
{
    allPersons(last: 2) {
        name
        age
    }
}
```

response:

```json
{
    "allPersons": [
        { "name": "Sarah", "age": 20 },
        { "name": "Alice", "age": 20 }
    ]
}
```

- complex, nested data:

```js
{
    allPersons {
        name
        posts {
            title
        }
    }
}
```

response:

```json
{
    "allPersons": [
        {
            "name": "Johnny",
            "posts": [
                { "title": "graphQL is awesome" },
                { "title": "Relay is a powerful GraphQL client" }
            ]
        }
        {
            "name": "Sarah",
            "posts": [
                { "title": "HHow to get started with React & GraphQL" }
            ]
        }
        /* ..... */
    ]
}
```

- making changes to the endpoint, Mutations!

**creating** new data

```js
// root fild is createPerson
    mutation {
        createPerson(name: "Bob", age: 36) {
            id
        }
    }
```

response:

```json
{
    "createPerson": {
        "id": 34
    }
}
```

**updating** new data

  - **realtime** updates, the tag `subscription` takes care of that. So it creates a stady connection to the server, for example:

```js
    subscription {
        newPerson {
            name
            age
        }
    }
```

This will create a connections to the server an listening to an event of when a new person is created

request:

....

**deleting** new data

**Creating a schema**:

based on the examples given on top we will design the most important part of graphQL, **THE SCHEMA**!:

```js
type Person {
    id: ID!
    name: String!
    age: Int!
    posts: [Post!]!
}

type Post {
    id: ID!
    title: String!
    author: Person!
}

// 3 root types allPersons, createPerson, newPerson
type Query {
    allPersons(last: Int): [Person!]!
    allPosts(last: Int): [Post!]!
}

type Mutation {
    createPerson(name: String!, age: String!): Person!
    updatePerson(id: ID!, name: String!, age: String!): Person!
    deletePerson(id: ID!): Person!
    createPost(title: String!): Post!
    updatePost(id: ID!, title: String!): Post!
    deletePost(id: ID!): Post!
}

type Subscription {
    newPerson: Person!
    updatedPerson: Person!
    deletedPerson: Person!
    newPost: Post!
    updatedPost: Post!
    deletedPost: Post!
}
```

so with this we can see that at the moment we can just use the persons and the post type is not being used :(

## **Architecture**

1 - GraphQL server with a connected database

2 - GraphQL server to integrate existing system
