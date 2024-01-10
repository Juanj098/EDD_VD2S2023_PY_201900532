import { NavBarT } from "../components/navbarT"
import tutorimg from '../assets/maestro.png'
export function Tutor(){
    return (
        <>
            <NavBarT/>
            <div className="container">
                <h1> Bienvenido Tutor...</h1>
                <img className="imgtutor" src={tutorimg} alt="" />
            </div>
        </>
    )
}