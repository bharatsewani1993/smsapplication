{{ template "main.tpl" . }}

{{ define "pagetitle" }}
  Send SMS
{{ end }}

{{ define "breadcrumbs" }}
    Send SMS
{{ end }}


{{ define "pagecontent" }}
<label for="phone_num">Phone Number:</label>
<input id="phone_num" class="form-control col-lg-3 col-md-6 vol-sm-6 col-xs-12" type="number" name="phone_num"/>
<label for="msg">Message:</label>
<textarea id="msg" class="form-control col-lg-6" name="msg" style="height:200px; resize:none;"></textarea>
<br/>
<button class="btn btn-success">Send Message</button>
<br/>
<br/>
<table id="smslogs" class="table table-hover table-responsive-lg table-striped">
<thead>
    <tr>
        <th>Sr</th>
        <th>Phone Number</th>
        <th>Message</th>
        <th>Date Time</th>
        <th>Action</th>
    </tr>
</thead>
<tbody>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Resend/Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Resend/Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Resend/Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Resend/Delete</td>
  </tr>
  <tr>
    <td>1</td>
    <td>1234567890</td>
    <td>Test</td>
    <td>2018-05-04 5:00:12 PM </td>
    <td>Resend/Delete</td>
  </tr>
</tbody>
</table>
{{ end }}

{{ define "js" }}
<script>
  var sendsms = function(){
    alert("SEND SMS CALLED");
  }

  $(document).ready(function() {
      $('#smslogs').DataTable();
  });
  
</script>
{{ end }}
