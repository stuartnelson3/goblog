$(document).on('click', '.preview', function() {
  var body = $('textarea').val();
  var $preview = $('.preview');
  var params = {
    url: '/preview',
    type: 'POST',
    data: {'post.Body': body},
    success: function(data) {
      $preview.html(data);
    }
  };
  $preview.empty();
  $.ajax(params);
});
