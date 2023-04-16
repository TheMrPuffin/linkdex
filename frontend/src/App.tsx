import { Container, Typography} from '@mui/material';
import Links from './components/Links'

function App(){
  return <Container>
    <Typography variant='h1' sx={{ my: 4, textAlign: 'center', color: 'primary.main'}}>LinkDex</Typography> 
    <Links />
  </Container>
}

export default App;