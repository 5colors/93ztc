{{define "navibar"}}
<header class="navbar navbar-transparent" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
        <span class="sr-only">
          导航菜单
        </span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a href="/" class="navbar-brand">
        <img src="/static/img/train2.png" style="width:70px;height:50px;"></img>
      </a>
    </div> <!--/.navbar-header -->
    <div class="navbar-collapse collapse">
      <ul class="nav navbar-nav">
        <li>
          {{if .IsLogined}}
          <a href="/dailyreport">
            <strong>店铺日报表</strong>
          </a>
          {{else}}
          <a href="/products">
            <strong>产品介绍</strong>
          </a>
          {{end}}
        </li>
        <li>
          {{if .IsLogined}}
          <a href="/monthlyreport">
            <strong>店铺月报表</strong>
          </a>
          {{else}}
          <a href="/service">
            <strong>服务体系</strong>
          </a>
          {{end}}
        </li>
        <li class="dropdown">
          <a href="/#" class="dropdown-toggle" data-toggle="dropdown">
            <strong>更多</strong> <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            {{if .IsLogined}}
            <li>
              <a href="/account#accinfo">
                账户信息管理
              </a>
            </li>
            {{end}}
            {{if .IsLogined}}
            <li>
              <a href="/service">
                服务体系管理
              </a>
            </li>
            {{end}}
            <li>
              <a href="/articles">
                电商赚钱之道
              </a>
            </li>
            <li>
              <a href="/faq">
                常见问题
              </a>
            </li>
            <li>
              <a href="/roadmap">
                技术路线图
              </a>
            </li>
            <li>
              <a href="/aboutus#partners">
                合作伙伴
              </a>
            </li>
            <li class="divider"></li>
            <li>
              <a href="/aboutus">
                关于我们
              </a>
            </li>            
          </ul>
        </li>
      </ul>

            <ul class="nav navbar-nav navbar-right">
              <li class="dropdown">
                <a href="/#" class="dropdown-toggle" data-toggle="dropdown">
                  <img src="/static/img/star.png" width="20" height="18">
                    <span class="caret"></span>
                </a>
                <ul class="dropdown-menu">
                  <li>
                    <p>
                      金冠直通车，只提供最专业的数据和运营服务，让中小电商也能享受大数据带来的运营便利.
            <img src="/static/img/qr.png"></img>&nbsp;&nbsp;微信扫一扫体验！
                    </p>
                  </li>
                  <li class="divider"></li>
                  <li>
                    <a href="/products">
                      马上了解更多
                    </a>
                  </li>
                </ul>
              </li>
              <li class="phone">    
          <strong>      
                  <abbr title="微信号">wx</abbr>: crownexp
          </strong>
              </li>              
              <li>
                  {{if .IsLogined}}                   
                  {{else}}
                    <a href="/login" class="btn btn-navbar-slide-blue">
                      登陆
                    </a>
                 {{end}}
              </li>
              <li>&nbsp;&nbsp;</li>
              <li>
                {{if .IsLogined}}
                <a href="/logout" class="btn btn-navbar-transparent">
                  安全退出
                </a>
                {{else}}
                <a href="/register#accreg" class="btn btn-navbar-transparent">
                  15秒马上加入！
                </a>
                {{end}}
              </li>
            </ul>
          </div> <!--/.nav-collapse -->
        </div> <!--/.container -->
      </header> <!--/.navbar-transparent -->
      <header class="navbar-slide" style="display: none;">
        <div class="navbar navbar-default" role="navigation">
          

  <div class="container">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
        <span class="sr-only">
          菜单
        </span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a href="/" class="navbar-brand">
        金冠·直通车
      </a>
    </div> <!--/.navbar-header -->
    <div class="navbar-collapse collapse">
      <ul class="nav navbar-nav">
        <li>
          {{if .IsLogined}}
            <a href="/dailyreport">
              <strong>店铺日报表</strong>
            </a>
          {{else}}
            <a href="/products/">
              <strong>产品介绍</strong>
            </a>
          {{end}}
        </li>
        <li>
          {{if .IsLogined}}
            <a href="/monthlyreport">
              <strong>店铺月报表</strong>
            </a>
          {{else}}
            <a href="/service">
              <strong>服务体系</strong>
            </a>
          {{end}}
        </li>
        <li class="dropdown">
          <a href="/#" class="dropdown-toggle" data-toggle="dropdown">
            <strong>更多</strong> <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            {{if .IsLogined}}
            <li>
              <a href="/account#accinfo">
                账户信息管理
              </a>
            </li>
            {{end}}
            {{if .IsLogined}}
            <li>
              <a href="/service">
                服务体系管理
              </a>
            </li>
            {{end}}
            <li>
              <a href="/articles">
                电商赚钱之道
              </a>
            </li>
            <li>
              <a href="/faq">
                常见问题
              </a>
            </li>
            <li>
              <a href="/roadmap">
                技术路线图
              </a>
            </li>
            <li>
              <a href="/aboutus#partners">
                合作伙伴
              </a>
            </li>
            <li class="divider"></li>
            <li>
              <a href="/aboutus">
                关于我们
              </a>
            </li>
          </ul>
        </li>
      </ul>
              <ul class="nav navbar-nav navbar-right">
                <li class="phone">                
                    <strong>+微信: crownexp</strong>
                </li>
                <li>
                  {{if .IsLogined}}                   
                  {{else}}
                    <a href="/login" class="btn btn-navbar-slide-blue">
                      登陆
                    </a>
                 {{end}}
                </li>
                <li>
                  {{if .IsLogined}}
                    <a href="/logout" class="btn btn-navbar-slide-plans">
                      安全退出
                    </a>
                  {{else}}
                    <a href="/register#accreg" class="btn btn-navbar-slide-plans">
                      15秒马上加入！
                    </a>
                 {{end}}
                </li>
              </ul>
            </div> <!--/.nav-collapse -->
          </div> <!--/.container -->
        </div> <!--/.navbar-default -->
      </header> <!--/.navbar-slide -->    

    {{end}}