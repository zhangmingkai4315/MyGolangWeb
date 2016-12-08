function login(){
     var email = $("#login_email").val()
     var password = $("#login_password").val()
     //we should do more client validation for users.but right now i will left this work for server.
     if (email.length==0 || password.length==0){
        return false;
     }
     return true;
}