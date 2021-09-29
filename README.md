# Microservices for web chat and event invites
This project uses a GraphQL to handle web chats and event CURD
### Model 
- User
- Event
### Seeder
object factory manufactures model objects

## Graphql
```yaml
type user{
    "name": str, 
    "password": str,
    "phone": str, required
    "email": str, 
    "emailConfirmed": bool, 
    "token": str,
    "createdAt": timestamp,
    "updatedAt": timestamp,
    "deletedAt": timestamp,
}
```

```yaml
type events{
    "name": str, required
    "isVirtual": bool, required
    "isPrivate": bool, required
    "size": int, required
    "owner": user_id, required
    "guests":[
      {
      "user_id": object_id, 
      "phone": str,
      "confirmed": nil},
      ], required
    "intro": str,
    "createdAt": timestamp,
    "updatedAt": timestamp,
    "deletedAt": timestamp,
    "startAt": timestamp, required
    "endAt": timestamp, required
    "location":[
      "zoom":{
        "id":str,
        "password":str},
      "address":{
          "street":str,
          "city":str,
          "state":str}], required
    "chats":[
      {
        "user_id": object_id,
        "name": str,
        "message": str},]
}
```

## Microservices
Microservices use RabbitMQ for internal communication.

- GraphQL for external communication
- Restful API for internal communication

### Event Service
#### API
authenticated post
```yaml
mutate event{}
```
response
```yaml
{
  "data":{
    "event":{
      "_id": object_id,
    }
  }
}
```
#### Test
- [ ] service runnable
- [ ] service stoppable
- [ ] read events
- [ ] create event
- [ ] update event
- [ ] delete event

### User Service
#### API
#### tests
- [ ] login to get token
- [ ] logout to delete token
- [ ] use token to get protected resources 
- [ ] use token to push protected resources 
- [ ] use token to delete protected resources 

### Invite Service

### Chat Service
#### API
authenticated post
```yaml
mutate events(id){
  "chats": push
      { 
      "user_id": object_id,
      "name": str,
      "message": str,
      }
}
```
response
```yaml
{
  "data":{
    "event":{
      "_id": object_id,
      "chats":[...]
    }
  }
}
```

