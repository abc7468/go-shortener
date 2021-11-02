const form = document.querySelector(".wrapper form"),
fullURL = form.querySelector("input"),
shortenBtn = form.querySelector("form button");
delBtns = document.querySelectorAll(".delbtn");

function deleteURL(id){
    var data = JSON.stringify({ origin_url: id })
    
    $.ajax({
        type:"DELETE",
        url:"/api/shortener",
        headers:{
            "Accept": "application/json",
            "Content-Type" : "application/json",
        },
        data : data,
        async: false,
        success:function(){
            window.location = '/'
        },
        error:function(request,status,error){
            
                alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
            
        }
    });
}




shortenBtn.onclick = ()=>{
    var data = JSON.stringify({ origin_url: fullURL.value })
    $.ajax({
        type:"POST",
        url:"/api/shortener",
        headers:{
            "Accept": "application/json",
            "Content-Type" : "application/json",
        },
        data : data,
        async : false,
        success:function(){
            window.location = '/'
        },
        error:function(request,status,error){
            
                alert("code:"+request.status+"\n"+"message:"+request.responseText+"\n"+"error:"+error);
            
        }
    });
}
