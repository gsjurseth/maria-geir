// import the `Kafka` instance from the kafkajs library
const { Kafka } = require("kafkajs")

const kafka = new Kafka({
  clientId: 'app',
  brokers: ['kafka.local:9092']
})

const producer = kafka.producer()

const topic = "messages";

let i = 1;
const run = async () => {
  // Producing
  await producer.connect()
  console.log("Connected and about to send messages")

  setInterval(async () => {
    try {
      // send a message to the configured topic with
      // the key and value formed from the current value of `i`
      await producer.send({
        topic,
        messages: [
          {
            key: `${i}`,
            value: JSON.stringify({ msg: "this is message " + i }),
          },
        ],
      })

      // if the message is written successfully, log it and increment `i`
      console.log("writes: ", i)
      i++
    } catch (err) {
        console.error("could not write message " + err)
    }
  }, 1000)
}
 
run().catch(console.error)
