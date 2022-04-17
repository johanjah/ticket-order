# Gin + Rel Ticket Order Backend

## Installation

### Prerequisite

1. Migrasi terlebih dahulu database dengan menjalankan file pada
   ```
   <project>/db/migrations/main.go
   ```

### Running

1. Jalankan file dari
   ```
   <project>/cmd/api/main.go
   ```

### Database design
File database design : https://github.com/johanjah/ticket-db-design

buka file db.drawio di situs draw.io

### API design
```
GET    /roles/   
POST   /roles/      
GET    /roles/:ID  
PATCH  /roles/:ID  
DELETE /roles/:ID  
DELETE /roles/     
GET    /users/         
POST   /users/         
GET    /users/:ID      
PATCH  /users/:ID      
DELETE /users/:ID      
DELETE /users/         
GET    /paymentTypes/  
POST   /paymentTypes/   
GET    /paymentTypes/:ID 
GET    /events/:ID       
PATCH  /events/:ID         
DELETE /events/:ID         
DELETE /events/            
GET    /shoppingCharts/    
POST   /shoppingCharts/    
GET    /shoppingCharts/:ID   
PATCH  /shoppingCharts/:ID   
DELETE /shoppingCharts/:ID   
DELETE /shoppingCharts/
```

### Contoh body json
/// create role
```
// role 1
{
   "role_name" : "Administrator"
}

// role 2
{
	"role_name" : "User"
}
```


/// create user
```
// user 1
{
   "username" : "admin",
   "password" : "admin",
   "email" : "admin@google.com",
   "first_name" : "admin first name",
   "last_name" : "admin last name",
   "role_id" : 1
}

// user 2
{
   "username" : "user",
   "password" : "user",
   "email" : "user@google.com",
   "first_name" : "user first name",
   "last_name" : "user last name",
   "role_id" : 2
}
```




/// create payment type
```
{
   "payment_name" : "BCA",
   "payment_description" : "BCA payment type"
}
```

/// create event
```
{
   "event_name" : "Seminar IT Indonesia",
   "event_description" : "Seminar IT Indonesia 2022",
   "event_location" : "Jakarta",
   "event_start_date" : "2022-04-17T22:09:26+07:00",
   "event_end_date" : "2022-04-18T22:09:26+07:00",
   "base_price" : 100000
}
```

/// shopping chart
```
{
   "event_id" : 1,
   "user_id" : 2,
   "ticket_count" : 5
}
```

### Credit
Kredit kode sumber : https://github.com/go-rel/gin-example/

Note : fitur payment belum ada 🙏