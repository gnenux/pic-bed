(function () {
    function addLoadListener(fn) {
		if (typeof window.addEventListener != "undefined") {
			window.addEventListener("load", fn, false)
		} else {
			if (typeof document.addEventListener) {
				document.addEventListener("load", fn, false)
			} else {
				if (typeof window.attachEvent != "undefined") {
					window.attachEvent("load", fn)
				} else {
					var oldfn = window.onload;
					if (typeof window.onload != "function") {
						window.onload = fn
					} else {
						window.onload = function() {
							oldfn();
							fn()
						}
					}
				}
			}
		}
    }
 
    addLoadListener(function(){
        function sendSignupData() {
            var data = new Object();

            var form = document.getElementById("form-signup");

            var fd = new FormData(form);
            for(var key of fd.keys()){
                data[key] = fd.get(key);
            }

            var xhr = new XMLHttpRequest();
            xhr.addEventListener("load", function(event){
                if(event.target.responseText!= "undefined" && event.target.responseText!=null) {
                    var res = JSON.parse(event.target.responseText);
                    if (res["error"]!= "undefined" && res["error"]!=null ){
                        alert(res["error"]);
                    }
                    else{
                        window.location.replace("/");
                    }
                }
            });

            xhr.addEventListener("error",function(event){
                alert("Oups! Something goes wrong.");
            })

            xhr.open("POST","http://127.0.0.1:8080/v1/user");
            xhr.send(JSON.stringify(data));
        }

        var form = document.getElementById("form-signup");
        form.addEventListener("submit", function(event){
            event.preventDefault();
            sendSignupData();
        })
    });

})();
  
function validatePassword() {
    var inputPassword = document.getElementById("inputPassword").value;
    var confirmPassword = document.getElementById("confirmPassword").value;
    if(inputPassword != confirmPassword ){
        document.getElementById("passwordMessage").innerHTML = "<font color='red'>两次密码不相同</font>";
        return false;
    }
    else{
        document.getElementById("passwordMessage").innerHTML = "";
    }
    return true;
}