const WebSocket = require('ws');
const { Kafka } = require("kafkajs")


const kafka = new Kafka({
  clientId: 'app',
  brokers: ['kafka.local:9092']
})

const wss = new WebSocket.Server({ port: 8000 });

console.log("Websocket server should be up and we've createad our kafka consumer")

const topic = "messages";
const run = async (o) => {
  console.log('####### The headers: %v', o.req.headers)
  const groupId = o.req.headers['x-group-id'] || 'app'
  const consumer = kafka.consumer({groupId})
  // Consuming
  await consumer.connect()
  await consumer.subscribe({ topic, fromBeginning: true })
 
  await consumer.run({
    eachMessage: async ({ t, partition, message }) => {
      console.log({
        partition,
        offset: message.offset,
        value: message.value.toString(),
      });
      //ws.send({'msg': message.value, 'topic': t});
      o.ws.send(message.value);
    },
  })
}
 

wss.on('connection', function connection(ws,req) { ws.on('message', function incoming(message) {
  console.log('received: %s', message);
  });
  run({ws,req}).catch(console.error)
});
