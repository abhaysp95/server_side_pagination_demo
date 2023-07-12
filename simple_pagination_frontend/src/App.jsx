import './App.css';
import { Container } from 'react-bootstrap';
import Header from './Header';
import CarSaleDetail from './CarSaleDetail';

function App() {
  return (
	  <>
		<Container>
			<Header />
			<CarSaleDetail />
		</Container>
	  </>
  );
}

export default App;
