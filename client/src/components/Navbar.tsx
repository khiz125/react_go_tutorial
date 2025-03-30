import { Box, Button, Container, Flex, Text } from '@chakra-ui/react';
import React, { SetStateAction } from 'react'
import { IoMoon } from 'react-icons/io5';
import { LuSun } from 'react-icons/lu';

type Props = {
  colorMode: "light" | "dark";
  setColorMode: React.Dispatch<SetStateAction<"light" | "dark">>
}
const Navbar: React.FC<Props> = ({ colorMode, setColorMode }) => {


  const toggleColorMode = () => {
    if (colorMode === "light") {
      setColorMode("dark")
    } else {
      setColorMode("light")
    }
  }

  return (
    <Container maxW={"900px"}>
      <Box bg={`${colorMode === "light" ? "#333" : "white"}`} transition="background 0.2s ease-in-out" px={4} my={4} borderRadius={5}>
        <Flex h={16} alignItems="center" justifyContent="space-between">
          <Flex
            justifyContent={"center"}
            alignItems={"center"}
            gap={3}
            display={{ base: "none", sm: "flex" }}
          >
            <Text color={`${colorMode === "light" ? "white" : "#333"}`} fontSize={"40"}>Go + React</Text>
            <Text color={`${colorMode === "light" ? "white" : "#333"}`} fontSize={"40"}>= Web applicaton</Text>
          </Flex>
          <Flex alignItems={"center"} gap={3}>
            <Text color={`${colorMode === "light" ? "white" : "#333"}`} fontSize={"lg"} fontWeight={500}>
              Daily Tasks
            </Text>
            {/* Toggle Color Mode */}
            <Button onClick={toggleColorMode} transition="background 0.2s ease-in-out">
              {colorMode === "light" ? <IoMoon /> : <LuSun size={20} />}
            </Button>
          </Flex>
        </Flex>
      </Box>
    </Container>

  )
}

export default Navbar;