document.addEventListener("DOMContentLoaded", function () {
  // Try direct initialization of individual date fields
  const startDate = document.querySelector('input[name="start"]');
  const endDate = document.querySelector('input[name="end"]');

  if (startDate) {
    new Datepicker(startDate, {
      format: "yyyy-mm-dd",
      autohide: true,
    });
  }

  if (endDate) {
    new Datepicker(endDate, {
      format: "yyyy-mm-dd",
      autohide: true,
    });
  }
});
