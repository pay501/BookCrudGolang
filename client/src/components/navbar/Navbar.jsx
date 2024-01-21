import style from "./navbar.module.css"
import NavbarLinkPage from "./navbarLink/NavbarLink"

export default function () {

    return(
        <div className={`${style.container} bg-neutral`}>
            <div className={`${style.logo}`}>
                JathaDev
            </div>
            <div>
                <NavbarLinkPage/>
            </div>
        </div>
    )
}