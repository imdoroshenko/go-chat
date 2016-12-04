const socket = new WebSocket("ws://192.168.1.132:3000")
socket.addEventListener('open', (e) => {
  console.log(e)
})

const messageTpl = document.createElement('div')

function main() {
  const form = document.querySelector('#form'),
        message = document.querySelector('#message'),
        feed = document.querySelector('#feed')

  form.addEventListener('submit', (e) => {
    console.log('submit', message.value)
    sendMessage(message)
    e.preventDefault()
  })

  socket.addEventListener('message', (e) => {
    console.log('message', e)
    renderMessage(feed, e.data)
  })
}

function renderMessage(container, data) {
  const messageNode = messageTpl.cloneNode(true),
        message = JSON.parse(data)
  messageNode.appendChild(document.createTextNode(message.message))
  container.appendChild(messageNode)
}

function sendMessage(message) {
  socket.send(JSON.stringify({
    message: message.value
  }))
  message.value = ''
}



