const btn = document.getElementById('btn')

const form = document.getElementById('myform');
unique = document.getElementById('unique').getAttribute('unique'),
email = document.getElementById('email').getAttribute('email'),
id = btn.getAttribute('data_id');

const verified=true

console.log(email+" "+unique+"p");
form.addEventListener("submit",function(event){
   event.preventDefault();

    data ={
      id:unique,
      email:email,
      verified:true,
      

    }
  

        sendToApi(data)
        
   
    
     

})
btn.addEventListener('click',()=>{
  alert(id)
})

function sendToApi(data){
  const URL = 'http://localhost:8000/voterverify'


// console.log(JSON.stringify(data))

return  fetch(URL,{
method:'PUT',
headers:{
  'Content-Type':'application/json'
},
body:JSON.stringify(data)

})
.then(response=>{
if(!response.ok){
    alert("Not verify check internet connetion")
  throw new Error('Network response was not ok ')
}
return response.text()
})
.then(data=>{

const ans = JSON.parse(data)
// console.log(ans);
alert(ans.message);
location.href="http://localhost:8000/verify"


})
.catch(error=>{
   
console.log('Fetch error :',error)
})
}

//for deleting a row

document.getElementById("logout1").addEventListener("click", function() {
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