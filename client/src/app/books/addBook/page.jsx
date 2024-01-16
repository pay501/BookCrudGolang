"use client"
import axios from "axios";
import { useState } from "react";
import Swal from "sweetalert2";
import style from "./addBook.module.css";

export default function AddBookPage(){
    
    const [form, setForm] = useState({
        name:"",
        author:"",
        description:"",
        price:0.0,
    })

    function inputChange(e){
        e.preventDefault();
        const {name, value} = e.target
        setForm((prevForm)=>({...prevForm, [name]:value}))
    }

    function addBook (e){
        e.preventDefault();
        axios.post("http://localhost:8080/addBook",{
            name:form.name,
            author:form.author,
            description:form.description,
            price:parseFloat(form.price)
        },{withCredentials:true})
        .then(res=>{
            console.log(res.data)
            if(res.data.message === "Add successful"){
                Swal.fire({
                    icon:"success",
                    title:"Book has added already!",
                    footer:"Thank you!"
                })
            }else{
                Swal.fire({
                    icon:"error",
                    title:"There is something wrong, can't add book.",
                    footer:"Thank you!"
                })
            }
        })
    }

    return(
        <div>
            <div className={style.container}>
                <form className="flex flex-col gap-3">
                    <div>
                        <input type="text" placeholder='Title' name='name' onChange={inputChange}/>
                    </div>
                    <div>
                        <input type="text" placeholder='Author' name='author' onChange={inputChange}/>
                    </div>
                    <div>
                        <input type="text" placeholder='Description' name='description' onChange={inputChange}/>
                    </div>
                    <div>
                        <input type="number" placeholder='Price' name='price' onChange={inputChange}/>
                    </div>
                    <button
                        className={`bg-blue-500 rounded-lg py-1 w-40`}
                        onClick={addBook}
                    >
                        Add
                    </button>
                </form>
            </div>
        </div>
    )
}