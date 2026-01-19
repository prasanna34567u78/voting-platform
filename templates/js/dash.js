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