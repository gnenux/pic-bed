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

            form.submit();
       }

       inputs["images"].click();
    })
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