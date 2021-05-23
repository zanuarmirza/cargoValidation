## Cargo Validation

### endpoint
- validate 
payload sample :
```
{
    "listItem":[
    {
        "id":"item001",
        "type":"L",
        "weigth":1,
        "temperature":18
    },
    {
        "id":"item002",
        "type":"S",
        "weigth":3,
        "temperature":25,
        "volume":{
            "height":2,
            "width":2,
            "length":2
        }
    },
    {
        "id":"item003",
        "type":"S",
        "weigth":10,
        "temperature":25,
        "volume":{
            "height":8,
            "width":2,
            "length":2
        }
    },
    {
        "id":"item004",
        "type":"S",
        "weigth":5,
        "temperature":25,
        "volume":{
            "height":8,
            "width":2,
            "length":2
        }
    }
]
}
```