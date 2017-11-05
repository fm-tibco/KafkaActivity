# tibco-kafkapub
This activity provides your flogo application the ability to send a Kafka message


## Installation

```bash
flogo install github.com/rickdantz/KafkaActivity
```

## Schema
Inputs and Outputs:

```json
{
 "inputs":[
    {
      "name": "Topic",
      "type": "string",
      "required": true
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    },
  ]
}
```
## Settings
| Setting    | Description                                                                                                                                                                                             | Cardinality |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------|
| BrokerUrls | The Kafka cluster to connect to                                                                                                                                                                         | Required    |
| Token      | The Kafka topic on which to place the message                                                                                                                                                           | Required    |
| Message    | The text message to send                                                                                                                                                                                | Required    |
| user       | If connectiong to a SASL enabled port, the userid to use for authentication                                                                                                                             | Optional    |
| password   | If connectiong to a SASL enabled port, the password to use for authentication                                                                                                                           | Optional    |
| truststore | If connectiong to a TLS secured port, the directory containing the certificates representing the trust chain for the connection.  This is usually just the CACert used to sign the server's certificate | Optional    |

## Outputs
| Value     | Description                                            |
|-----------|--------------------------------------------------------|
| partition | Documents the partition that the message was placed on |
| offset    | Documents the offset for the message                   |

```
