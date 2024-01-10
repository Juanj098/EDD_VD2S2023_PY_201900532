import { NavBar } from "../components/navbar"
import ad from '../assets/blogger.png'

export function Admin(){
    return(
        <>
            <NavBar/>
            <div className="container">
                <h1>Bienvenido Administrador...</h1>
                 <img src={ad} alt="" />
                
            </div>
        </>
                
    )
}