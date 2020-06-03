
(function($) {
    $("#zakaat_in_btc").hide();
    $(".rest").hide();
    $(".zkcss").hide();

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

            // var btc = Math.round(result * 100) / 100;
            var obj = JSON.parse(result);
            
            var btc = Math.round(obj.zakaat_in_bitcoin* 100) / 100;
            
            var usd = Math.round(obj.zakaat_in_USD* 100) / 100;
            var gold = Math.round(obj.zakaat_in_gold* 100) / 100;
            var rm = Math.round(obj.zakaat_in_rm* 100) / 100;
            console.log(btc);
            if (btc < 0){
                $("#welcome_msg").hide();
                $(".zkcss").show();
                $("#zakaat_in_btc").show();
                $("#zakaat_in_btc").html(0.00+ " BTC");

                $(".rest").show();
                $(".rest").html("<span>"+0.00 +" in USD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+0.00 +"gm in GOLD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+0.00 +" in MYR</span>");
            } else {
                $("#welcome_msg").hide();
                $(".zkcss").show();
                $("#zakaat_in_btc").show();
                $("#zakaat_in_btc").html(btc+ " BTC");

                $(".rest").show();
                $(".rest").html("<span>"+usd +" in USD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+gold +"gm in GOLD</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span>"+rm +" in MYR</span>");
            }

    })
})    


$(function() {
  $('a[href*=#]').on('click', function(e) {
    e.preventDefault();
    $('html, body').animate({ scrollTop: $($(this).attr('href')).offset().top}, 600, 'linear');
  });
});