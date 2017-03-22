(function(){
  var newscript = document.createElement('script');
  newscript.type = 'text/javascript';
  newscript.async = true;
  newscript.src = 'https://ajax.googleapis.com/ajax/libs/jquery/1.6.1/jquery.min.js';
  (document.getElementsByTagName('head')[0]||document.getElementsByTagName('body')[0]).appendChild(newscript);
})();

var module = (function() {
  var count = 0;
  return {
    increment: function() {
      count++;
    },
    show: function(value) {
      console.log(count);
      console.log($(value.childNodes[0].firstElementChild).text());
      console.log($(value.childNodes[1].firstElementChild).text());
      console.log($(value.childNodes[3].firstElementChild).text());
      console.log($(value.childNodes[4].firstElementChild).text());
    }
  };
})();

function showData(index,value){
  if(index==0){
  	var timestamp = value.firstElementChild;
    console.log(timestamp);
  } else if(index==1){
   return;
  } else {
    module.show(value);
    module.increment();
  }  
}

var trs = $("table")[2].firstElementChild.childNodes; //array

$.each(trs,function(index,value){ showData(index,value) });
