// eslint-disable-next-line no-unused-vars
import React from 'react'

function App() {
  return (
    <div>
        <h1>Hola, prueba 4</h1>
        {/* Comunicando el front con el back
        Que cuando le de click haga una peticion fetch
         */}
        <button onClick={async () => {
            {/* Esto va a ser asincrono */}
           const response = await fetch('http://localhost:3000/users')
           const data = await response.json()
           console.log(data)
        }
        }>oprimir</button>
    </div>
  )
}

export default App
