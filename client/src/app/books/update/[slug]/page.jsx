"use client"
import axios from "axios"
import Image from "next/image"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import Swal from "sweetalert2"
import style from "./updateBook.module.css"

export default function UpdateBookPage({ params }) {

    const router = useRouter()
    const id = params.slug
    const [form, setForm] = useState({
        name:"",
        author:"",
        description:"",
        price:0.0,
        image:""
    })

    async function getBook() {
        await axios.get(`http://localhost:8080/book/${id}`, { withCredentials: true })
            .then(res => {
                if (res.data.message === "return") {
                    router.push("/login")
                }else{
                    setForm({
                        name:res.data.result.name,
                        author: res.data.result.author,
                        description:res.data.result.description,
                        price: res.data.result.price,
                        image:res.data.result.image
                    })
                }
            })
    }

    function inputChange(e) {
        const {name, value} = e.target
        setForm((prevForm)=>({...prevForm, [name]:value}))
    }

    async function updateBook (e) {
        try{
            e.preventDefault();
            if(
                form.author === "" ||
                form.name === "" ||
                form.price === 0.0 ||
                form.description === "" ||
                form.image === ""
            ){
                Swal.fire({
                    icon: "error",
                    title: "Oops! Can't add book!",
                    text: "Please ,fill in all data"
                })
            }else{
                await axios.get("http://localhost:8080/checkSession",{withCredentials:true})
                    .then(res=>{
                        if(res.data.message === "return"){
                            console.log(res.data.message)
                            router.push("/login")
                        }else{
                            axios.put(`http://localhost:8080/updateBook/${id}`, {
                                name:form.name,
                                author:form.author,
                                description: form.description,
                                price: parseFloat(form.price),
                                image:form.image
                            },
                            {withCredentials:true})
                            .then(res=>{
                                if(res.data.message === "Update successful"){
                                    Swal.fire({
                                        icon:"success",
                                        title:"Update book successful",
                                        footer:"Thank you!",
                                        timer:"1500"
                                    })
                                    router.push("/manage/book")
                                }else{
                                    Swal.fire({
                                        icon:"error",
                                        title:"Update book failed!",
                                        text:"There is something wrong!"
                                    })
                                }
                            })
                        }
                    })
            }
        }catch(err){
            console.log(err)
        }
    }

    useEffect(() => {
        getBook()
    }, [])

    return (
        <div className={`${style.container}`}>
            <div className={`${style.imgContainer}`}>
                <Image
                    src= {`${form.image}`}
                    className={style.img}
                    alt=""
                    fill
                >
                </Image>
            </div>
            <div className={`${style.formContainer}`}>
                <div>
                    <h1 className="text-4xl font-bold">Welcome to the Library!</h1>
                    <p
                        className="text-lg font-semibold mt-4"
                    >
                        I am so appreciate that you were join our library.
                        <br /><span>
                            Wanna update your book?
                        </span>
                    </p>

                </div>
                <div>
                    <form className="flex flex-col gap-4 my-8">
                        <div>
                            <input
                                type="text"
                                placeholder='Title'
                                name='name'
                                className="input input-primary w-full max-w-md"
                                onChange={inputChange}
                                value={form.name}
                                required
                            />
                        </div>
                        <div>
                            <input
                                type="text"
                                placeholder='Author'
                                name='author'
                                className="input input-primary w-full max-w-md"
                                onChange={inputChange}
                                value={form.author}
                                required
                            />
                        </div>
                        <div>
                            <textarea
                                type="text"
                                placeholder='Description'
                                name='description'
                                className="input input-primary w-full max-w-md h-fit"
                                onChange={inputChange}
                                value={form.description}
                                cols="5"
                                rows="4"
                                required
                            />
                        </div>
                        <div>
                            <input
                                type="number"
                                placeholder='Price'
                                name='price'
                                className="input input-primary w-full max-w-md"
                                onChange={inputChange}
                                value={form.price}
                                required
                            />
                        </div>
                        <div>
                            <input
                                type="text"
                                placeholder='Image URL'
                                name='image'
                                className="input input-primary w-full max-w-md"
                                onChange={inputChange}
                                value={form.image}
                                required
                            /><br />
                            <label className="text-sm text-red-500">*Sorry for inconvenience*</label>
                        </div>
                        <button
                            className={`bg-blue-500 rounded-lg py-1 w-full max-w-md text-white font-semibold h-10`}
                            onClick={updateBook}
                        >
                            update
                        </button>
                    </form>
                </div>
            </div>
        </div>
    )
}