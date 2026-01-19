const add = document.getElementById('new')
const dialog = document.getElementById("dialog");
const closeButton = document.getElementById("closeButton");


add.addEventListener("click",()=> {
    dialog.style.display="block"
})
closeButton.addEventListener("click",()=> {
    dialog.style.display="none"
})

const form = document.getElementById('myform');
const position = document.getElementById('position'),
priority = document.getElementById('priority');
const id = document.getElementById('ID')



form.addEventListener("submit",function(event){
   event.preventDefault();

    data ={
       name:position.value,
       priority:parseInt(priority.value)

    }
  

        sendToApi(data)
        position.value=""
        priority.value=""
    
    dialog.style.display="none"
    
     

})

function sendToApi(data){
  const URL = 'http://localhost:8000/api/position'


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
    alert("Position name alread exists")
  throw new Error('Network response was not ok ')
}
return response.text()
})
.then(data=>{

const ans = JSON.parse(data)
// console.log(ans);
alert(ans.message);
location.href="http://localhost:8000/position"


})
.catch(error=>{
   
console.log('Fetch error :',error)
})
}

//for deleting a row

document.querySelectorAll('.delete_btn').forEach(button => {
  button.addEventListener('click',async function(){
    const row = this.getAttribute('data_id');
  
    const confirmation  = confirm("Are you sure you want to delete this record")
    if(confirmation){
      const response = await fetch(`/api/position_delete/${row}`,{
        method:'DELETE',
      });
      if(response.status == 200){
        document.getElementById(row).remove();
      }else{
        alert('Failed to delete')
      }
    }
  })
});

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