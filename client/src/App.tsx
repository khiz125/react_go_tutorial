import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import { useState } from 'react'

function App() {
  const [colorMode, setColorMode] = useState<"light" | "dark">("light")
  return (
    <Stack bg={`${colorMode === "light" ? "white" : '#333'}`} transition="background 0.2s ease-in-out" h="100vh">
      <Navbar colorMode={colorMode} setColorMode={setColorMode} />
      <Container>
        {/* <TodoForm />
        <TodoList /> */}
      </Container>
    </Stack>
  )
}

export default App
