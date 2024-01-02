import {Routes, Route} from 'react-router-dom'
import {Login} from './pages/login'
import {Admin} from './pages/Admin'
import {Tutor} from './pages/Tutor'
import {Estudiante} from './pages/Estudiante'
import './index.css'
export function App(){
    return(
        <div className='aplicacion'>
            <Routes>
                <Route path='/' element={<Login/>}/>
                <Route path='/Admin' element={<Admin/>}/>
                <Route path='/Tutor' element={<Tutor/>}/>
                <Route path='/Estudiante' element={<Estudiante/>}/>
            </Routes>
        </div>
    )
}