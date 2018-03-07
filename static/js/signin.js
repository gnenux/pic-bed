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

  addLoadListener(function () {

    function sendData() {

    var loginInfo = new Object();

    var form = this.document.getElementById("form-signin")

    var fd = new FormData(form);
    for(var key of fd.keys()){
        loginInfo[key] = fd.get(key);
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


      xhr.addEventListener("error", function(event) {
        alert('Oups! Something goes wrong.');
      });
  
      xhr.open("POST", "http://127.0.0.1:8080/v1/user/login");

      xhr.send(JSON.stringify(loginInfo));
    }

    var form = document.getElementById("form-signin");

    form.addEventListener("submit", function (event) {
      event.preventDefault();
  
      sendData();
    });
  });
  
})();