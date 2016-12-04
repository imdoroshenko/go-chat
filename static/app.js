const socket = new WebSocket('ws://localhost:3000')
socket.addEventListener('open', (e) => {
  console.log(e)
})

const messageTpl = document.createElement('div')

function main() {
  const form = document.querySelector('#form'),
        message = document.querySelector('#message'),
        feed = document.querySelector('#feed'),
        channel = document.querySelector('#channel')

  channel.addEventListener('change', (e) => {
    changeChannel(feed, channel)
  })

  form.addEventListener('submit', (e) => {
    console.log('submit', message.value, channel.value)
    sendMessage(message, channel)
    e.preventDefault()
  })

  socket.addEventListener('message', (e) => {
    console.log('message', e)
    renderMessage(feed, e.data)
  })
}

function changeChannel(feed, channel) {
  while(feed.firstChild){
    feed.removeChild(feed.firstChild)
  }
  socket.send(JSON.stringify({
    type:    'join',
    channel: channel.value
  }))
}

function renderMessage(container, data) {
  const messageNode = messageTpl.cloneNode(true),
        message = JSON.parse(data)
  switch (message.type) {
    case 'msg':
      messageNode.appendChild(document.createTextNode(message.value))
      container.appendChild(messageNode)
      break
    case 'join':
      messageNode.appendChild(document.createTextNode("Channel changed to " + message.channel))
      container.appendChild(messageNode)
      break
  }
  container.scrollTop = container.scrollHeight
}

function sendMessage(message, channel) {
  socket.send(JSON.stringify({
    type:    'msg',
    channel: channel.value,
    value:   message.value
  }))
  message.value = ''
}



