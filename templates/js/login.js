const form = document.getElementById('myform');
const email = document.getElementById('email'),
password = document.getElementById('password'),
role = document.getElementById('role');
console.log(role.value);
form.addEventListener("submit",function(event){
   event.preventDefault();

    data ={
        email:email.value,
        password:password.value
    }
  
   if(role.value === "ADMIN"){
    loginAdmin(data)
    email.value=""
password.value=""
   }else{
      loginVoter(data)
      email.value=""
      password.value=""
   }
       
  
    
     

})

function loginAdmin(data){
  const URL = 'http://localhost:8000/api/login'


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
    alert("user Email alread register or password is wrong")
  throw new Error('Network response was not ok ')
}
return response.text()
})
.then(data=>{

const ans = JSON.parse(data)
// console.log(ans);
alert(ans.message);
location.href="http://localhost:8000/dash"


})
.catch(error=>{
   
console.log('Fetch error :',error)
})
}
function loginVoter(data){
  const URL = 'http://localhost:8000/voter/login'


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
    alert("user Email alread register or password is wrong")
  throw new Error('Network response was not ok ')
}
return response.text()
})
.then(data=>{

const ans = JSON.parse(data)
// console.log(ans);
alert(ans.message);
location.href="http://localhost:8000/home"


})
.catch(error=>{
   
console.log('Fetch error :',error)
})
}



