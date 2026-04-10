
document.addEventListener('DOMContentLoaded', function () {
  
  var confirmBtn = document.querySelector('.btn--confirm');
  if (confirmBtn) {
    confirmBtn.addEventListener('click', function (e) {
      e.preventDefault();
      window.location.href = '/accepted';
    });
  }

  var declineBtn = document.querySelector('.btn--decline');
  if (declineBtn) {
    declineBtn.addEventListener('click', function (e) {
      e.preventDefault();
      window.location.href = '/denied';
    });
  }
});
