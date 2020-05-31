(function($) {

    var form = $("#signup-form");
    form.steps({
        headerTag: "h3",
        bodyTag: "fieldset",
        transitionEffect: "fade",
        labels: {
            previous : 'Previous',
            next : 'Submit',
            finish : 'Submit',
            current : ''
        },
        titleTemplate : '<div class="title"><span class="title-text">#title#</span><span class="title-number">0#index#</span></div>',
        onFinished: function (event, currentIndex)
        {
            console.log('Sumited');
        }
    });
    $("li").removeClass("disabled");    
    
})(jQuery);


var typed3 = new Typed('#typed', {
        strings: ['<i><b> Zakat</b></i>', '<i><b> Hajj</b></i>', '<i><b> Qurbani</b></i>', '<i><b> Fitra</b></i>'],
        typeSpeed: 100,
        backSpeed: 100,
        smartBackspace: true, // this is a default
        loop: true
      });

      function myFunction() {
      var x = document.getElementById("myTopnav");
      if (x.className === "topnav") {
        x.className += " responsive";
      } else {
        x.className = "topnav";
      }
    }



$("#zk_btn").click(function(){
    var amt = $("#amount").val();
    var asset = $("#asset").val();
    console.log(amt);
    console.log(asset);
    $.ajax({
        type: 'POST',
        url: '/zakaat',
        data: {amt:amt, asset: asset},
        
        }).done(function(result) { //use this

            console.log(result)
            // status = result
                    

    })
})    


$(function() {
  $('a[href*=#]').on('click', function(e) {
    e.preventDefault();
    $('html, body').animate({ scrollTop: $($(this).attr('href')).offset().top}, 600, 'linear');
  });
});