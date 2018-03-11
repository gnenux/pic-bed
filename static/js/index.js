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
    var form = document.getElementById("form-upload");
    var uploadbtn = document.getElementById("form-upload-btn");
    uploadbtn.addEventListener("click",function(event){
       var inputs = form.elements;

       inputs["images"].onchange = function (event) {
            var imageFiles = inputs["images"].files;
            console.log(imageFiles.length);
            for (let index = 0; index < imageFiles.length; index++) {
                var imageFile = imageFiles[index];
                console.log(imageFile.name);
                if(!validateFile(imageFile.name)){
                    alert(imageFile.name + " is not a image!");
                    return
                }
            }

            var fd = new FormData(form);
            var xhr = new XMLHttpRequest();
            xhr.addEventListener("load", function(event){
                if(event.target.responseText!= "undefined" && event.target.responseText!=null) {
                    var res = JSON.parse(event.target.responseText);
                    if (res["error"]!= "undefined" && res["error"]!=null ){
                        alert(res["error"]);
                    }
                }
            });

            xhr.addEventListener("error",function (event) {
                alert('Oups! Something goes wrong.');
            })

            xhr.open("POST",form.action);
            xhr.send(fd);
       }

       inputs["images"].click();
    })
  });
  
  addLoadListener(()=>{
      var album = document.getElementById("album");
      album.onclick = (e)=> {
          var classList = e.target.classList;
            album.parentElement
          if(classList.contains("btn-danger")){
             var card =  e.target.parentElement.parentElement.parentElement.parentElement.parentElement;
             card.parentElement.removeChild(card);          
          }
      }
  });

  function validateFile(fileName) {
      var index = fileName.lastIndexOf('.');
      if(index < 0){
          return false;
      }
      var name = fileName.toLowerCase();
      var suffix = name.substring(index+1);
      switch (suffix) {
          case "jpg":
          case "png":
          case "gif":
            return true;
          default:
          return false;
      }
  }

})();


function deleteImage(url,imageId) {
    var xhr = new XMLHttpRequest();
    xhr.addEventListener("load", function(event){
        if(event.target.status!="undefined" && event.target.status != null && event.target.status != 204){
            if(event.target.responseText!= "undefined" && event.target.responseText!=null) {
                var res = JSON.parse(event.target.responseText);
                if (res["error"]!= "undefined" && res["error"]!=null ){
                    alert(res["error"]);
                }
            }
        }else{
            event.target.style.display = "none";
        }
    });

    xhr.addEventListener("error",function (event) {
        alert('Oups! Something goes wrong.');
    })

    xhr.open("DELETE",url+"/"+imageId);
    xhr.send();
}