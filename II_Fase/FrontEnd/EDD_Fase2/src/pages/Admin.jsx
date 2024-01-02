import { Link } from "react-router-dom"
import imgB from '../assets/boton-de-encendido-apagado.png';
export function Admin(){
    return(
        <>
            <div className="container">
                <header className="hder-admin">
                    <Link to='/'><img className="btn-img" src={imgB} alt="" /></Link>
                </header>
                <main className="bdy-admin">
                    <h1>holaa..</h1>
                </main>
            </div>
        </>
                
    )
}