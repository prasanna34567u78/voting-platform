const form = document.getElementById('myform');
const email = document.getElementById('email'),
password = document.getElementById('password'),
confirm_password = document.getElementById('confirm_password')
form.addEventListener("submit",function(event){
   event.preventDefault();

    data ={
        email:email.value,
        password:password.value
    }
    if(password.value !== confirm_password.value){
        alert("passwords not matching");
    }else{

        sendToApi(data)
        email.value=""
    password.value=""
    confirm_password.value=""
    }
    // console.log(ans);
    
     

})

  function sendToApi(data){
  const URL = 'http://localhost:8000/add'


// console.log(JSON.stringify(data))

return  fetch(URL,{
method:'POST',
headers:{
  'Content-Type':'application/json'
},
body:JSON.stringify(data)

})
.then(response=>{
if(!response.ok){
    alert("user already exits GO to login page")
  throw new Error('Network response was not ok ')
}
return response.text();
})
.then(data=>{

const ans = JSON.parse(data)
console.log(ans);
alert(ans.message);
location.href="http://localhost:8000/login"


})
.catch(error=>{
   
console.log('Fetch error :',error)
})
}



document.getElementById("logout").addEventListener("click", function() {
  const URL = 'http://localhost:8000/logout'
  fetch(URL,{
    method:"POST",
  }).then(response=>{
    if(!response.ok){
        alert("Email alread exists")
      throw new Error('Network response was not ok ')
    }
    return response.text()
    }).then(data=>{
    const ans = JSON.parse(data)
    alert(ans.message)
    location.href="http://localhost:8000/login"
  })
})