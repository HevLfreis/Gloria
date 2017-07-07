package main

const (
	STATUS_SUCCESS     = 0
	STATUS_ERR         = -1
	STATUS_UNSUPPORTED = -5
	MONGO_DBNAME       = "gloria"
	BUS_TEMPLATE       = `<!DOCTYPE html><html lang="en">
						<head>
						<meta charset="UTF-8">
						<meta http-equiv="X-UA-Compatible" content="IE=edge">
    					<meta name="viewport" content="width=device-width, initial-scale=1">
    					<title>回家！回家！</title>
    					<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"/>
    					<link type="text/css" rel="stylesheet" href="https://cdn.bootcss.com/font-awesome/4.7.0/css/font-awesome.min.css"/>
    					</head>
    					<body>
    					<nav class="navbar navbar-default">
						<div class="container">
						<div class="navbar-header">
						<a class="navbar-brand" href="#">浦东11路</a>
						</div>
						</div>
						</nav>
						<div class="container">
						<ul class="list-group">
						<li class="list-group-item">
						<h2><i class="fa fa-bus"></i> 去金科路</h2><h3>还有%s站, %d分钟</h3>
						</li>
						<li class="list-group-item">
						<h2><i class="fa fa-bus"></i> 金科路回家</h2><h3>还有%s站, %d分钟</h3>
						</li>
						<li class="list-group-item">
						<h2><i class="fa fa-bus"></i> 去广兰路地铁站</h2><h3>还有%s站, %d分钟</h3>
						</li>
						<li class="list-group-item">
						<h2><i class="fa fa-bus"></i> 广兰路地铁站回家</h2><h3>还有%s站, %d分钟</h3>
						</li>
						</ul>
						<br><h4> Powered by 喜欢你的我呀 <i class="fa fa-heart"></i></h4><br>
						</div>
						</body>
						</html>`
)
