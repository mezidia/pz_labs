package engine

type Command interface { 
  Execute() 
} 

type Handler interface { 
  Post() 
}

