function login(){
     var email = $("#login_email").val()
     var password = $("#login_password").val()
     //we should do more client validation for users.but right now i will left this work for server.
     if (email.length==0 || password.length==0){
        return false;
     }
     return true;
}

function register(){
     var email = $("#register_email").val()
     var password = $("#register_password").val()
     var password2 = $("#register_password2").val()
     var username = $("#register_username").val()
     if(password!==password2){
        $("#register_password2").parent().addClass('is-invalid');
        return false;
     }
     //we should do more client validation for users.but right now i will left this work for server.
     if (email.length==0 || password.length==0||username.length==0){
        return false;
     }
     return true;
}

$(document).ready(function(){
    var close = document.getElementsByClassName("closebtn");
    for (var i = 0; i < close.length; i++) {
        close[i].onclick = function(){
            var div = this.parentElement;
            div.style.opacity = "0";
            setTimeout(function(){ div.style.display = "none"; }, 600);
        }
    }
})