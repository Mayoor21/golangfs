import CarView from "./Cars/CarsVieww"
import CarList from "./Cars/CarsList"
import CarCreate from "./Cars/CarsCreate"
import {BrowserRouter,Routes,Route} from "react-router-dom";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
            <Route path="" element={<CarList/>}/>
            <Route path="/car/list" element={<CarList/>}/>
            <Route path="/car/create" element={<CarCreate/>}/>
            <Route path="/car/view" element={<CarView/>}/>
        </Routes>
      
      </BrowserRouter>
    </>
  )
}

export default App