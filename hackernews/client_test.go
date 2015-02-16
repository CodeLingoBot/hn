package hackernews

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const TestResponse = `<html><head><meta name="referrer" content="origin"></meta><link rel="stylesheet" type="text/css" href="news.css?RhgQ45HmurMxjyZlSb9P"></link><link rel="shortcut icon" href="favicon.ico"></link><link rel="alternate" type="application/rss+xml" title="RSS" href="rss"></link><script type="text/javascript">
function byId(id) {
  return document.getElementById(id);
}

function vote(node) {
  var v = node.id.split(/_/);   // {'up', '123'}
  var item = v[1];

  // hide arrows
  byId('up_'   + item).style.visibility = 'hidden';
  byId('down_' + item).style.visibility = 'hidden';

  // ping server
  var ping = new Image();
  ping.src = node.href;

  return false; // cancel browser nav
} </script><title>Hacker News</title></head><body><center><table id="hnmain" op="news" border="0" cellpadding="0" cellspacing="0" width="85%" bgcolor="#f6f6ef"><tr><td bgcolor="#ff6600"><table border="0" cellpadding="0" cellspacing="0" width="100%" style="padding:2px"><tr><td style="width:18px;padding-right:4px"><a href="http://www.ycombinator.com"><img src="y18.gif" width="18" height="18" style="border:1px #ffffff solid;"></img></a></td><td style="line-height:12pt; height:10px;"><span class="pagetop"><b><a href="news">Hacker News</a></b><img src="s.gif" height="1" width="10"><a href="newest">new</a> | <a href="newcomments">comments</a> | <a href="show">show</a> | <a href="ask">ask</a> | <a href="jobs">jobs</a> | <a href="submit">submit</a></span></td><td style="text-align:right;padding-right:4px;"><span class="pagetop"><a href="login?goto=news">login</a></span></td></tr></table></td></tr><tr style="height:10px"></tr><tr><td><table border="0" cellpadding="0" cellspacing="0"><tr><td align="right" valign="top" class="title"><span class="rank">1.</span></td><td><center><a id="up_9050970" href="vote?for=9050970&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050970"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://websocketd.com/">Show HN: A Unixy approach to WebSockets</a><span class="sitebit comhead"> (websocketd.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050970">217 points</span> by <a href="user?id=joewalnes">joewalnes</a> <a href="item?id=9050970">6 hours ago</a>  | <a href="item?id=9050970">47 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">2.</span></td><td><center><a id="up_9051246" href="vote?for=9051246&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051246"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://imgur.com/gallery/jdNA6">Now, I can see wifi signals</a><span class="sitebit comhead"> (imgur.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051246">331 points</span> by <a href="user?id=bane">bane</a> <a href="item?id=9051246">4 hours ago</a>  | <a href="item?id=9051246">34 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">3.</span></td><td><center><a id="up_9051645" href="vote?for=9051645&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051645"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://github.com/lukehoban/es6features">A summary of ECMAScript 6 features</a><span class="sitebit comhead"> (github.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051645">11 points</span> by <a href="user?id=sevko">sevko</a> <a href="item?id=9051645">1 hour ago</a>  | <a href="item?id=9051645">3 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">4.</span></td><td><center><a id="up_9050793" href="vote?for=9050793&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050793"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://talks.golang.org/2015/state-of-go.slide#1">The State of Go</a><span class="sitebit comhead"> (golang.org)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050793">131 points</span> by <a href="user?id=xkarga00">xkarga00</a> <a href="item?id=9050793">7 hours ago</a>  | <a href="item?id=9050793">185 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">5.</span></td><td><center><a id="up_9050436" href="vote?for=9050436&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050436"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.nytimes.com/2015/02/15/world/bank-hackers-steal-millions-via-malware.html">Bank Hackers Steal Millions via Malware</a><span class="sitebit comhead"> (nytimes.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050436">149 points</span> by <a href="user?id=youlweb">youlweb</a> <a href="item?id=9050436">9 hours ago</a>  | <a href="item?id=9050436">47 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">6.</span></td><td><center><a id="up_9050666" href="vote?for=9050666&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050666"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.nytimes.com/2015/02/15/business/behind-monopoly-an-inventor-who-didnt-pass-go.html">Monopoly’s Inventor: The Progressive Who Didn’t Pass ‘Go’</a><span class="sitebit comhead"> (nytimes.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050666">28 points</span> by <a href="user?id=wallflower">wallflower</a> <a href="item?id=9050666">8 hours ago</a>  | <a href="item?id=9050666">1 comment</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">7.</span></td><td><center><a id="up_9051562" href="vote?for=9051562&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051562"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://volatileread.com/Wiki/Index?id=1087">C# 6 Exception Filters and How they are much more than Syntactic Sugar</a><span class="sitebit comhead"> (volatileread.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051562">13 points</span> by <a href="user?id=bursurk">bursurk</a> <a href="item?id=9051562">2 hours ago</a>  | <a href="item?id=9051562">3 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">8.</span></td><td><center><a id="up_9050646" href="vote?for=9050646&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050646"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://docs.racket-lang.org/style/index.html">How to Program Racket</a><span class="sitebit comhead"> (racket-lang.org)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050646">25 points</span> by <a href="user?id=shawndumas">shawndumas</a> <a href="item?id=9050646">8 hours ago</a>  | <a href="item?id=9050646">2 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">9.</span></td><td></td><td class="title"><span class="deadmark"></span><a href="https://segment.com/jobs/descriptions/infrastructure-engineer/" rel="nofollow">Segment is hiring engineers to revamp our infrastructure</a><span class="sitebit comhead"> (segment.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext">22 minutes ago</td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">10.</span></td><td><center><a id="up_9050480" href="vote?for=9050480&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050480"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://github.com/eugenkiss/7guis/wiki">7GUIs – A Notational Usability Benchmark for GUI Programming</a><span class="sitebit comhead"> (github.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050480">103 points</span> by <a href="user?id=desdiv">desdiv</a> <a href="item?id=9050480">9 hours ago</a>  | <a href="item?id=9050480">12 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">11.</span></td><td><center><a id="up_9050597" href="vote?for=9050597&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050597"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.nytimes.com/1988/02/25/us/brain-wound-eliminates-man-s-mental-illness.html">Brain Wound Eliminates Man's Mental Illness</a><span class="sitebit comhead"> (nytimes.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050597">24 points</span> by <a href="user?id=denysonique">denysonique</a> <a href="item?id=9050597">8 hours ago</a>  | <a href="item?id=9050597">15 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">12.</span></td><td><center><a id="up_9051288" href="vote?for=9051288&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051288"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.theguardian.com/higher-education-network/gallery/2015/feb/12/philosophical-transactions-of-the-royal-society-350-years-of-science-publishing-in-pictures">350 years of publishing from the world's oldest science journal in pictures</a><span class="sitebit comhead"> (theguardian.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051288">22 points</span> by <a href="user?id=Thevet">Thevet</a> <a href="item?id=9051288">4 hours ago</a>  | <a href="item?id=9051288">discuss</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">13.</span></td><td><center><a id="up_9049467" href="vote?for=9049467&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049467"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://spyder.wordpress.com/2014/03/16/why-ocaml-why-now/">Why OCaml, why now? (2014)</a><span class="sitebit comhead"> (spyder.wordpress.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049467">146 points</span> by <a href="user?id=lelf">lelf</a> <a href="item?id=9049467">16 hours ago</a>  | <a href="item?id=9049467">94 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">14.</span></td><td><center><a id="up_9051373" href="vote?for=9051373&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051373"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://amibeingtracked.com">Am I Being Tracked . com</a><span class="sitebit comhead"> (amibeingtracked.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051373">15 points</span> by <a href="user?id=MilnerRoute">MilnerRoute</a> <a href="item?id=9051373">3 hours ago</a>  | <a href="item?id=9051373">7 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">15.</span></td><td><center><a id="up_9050332" href="vote?for=9050332&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050332"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://medium.com/@DanielleMorrill/working-with-brad-4f79d8859443">Post Series A Life</a><span class="sitebit comhead"> (medium.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050332">68 points</span> by <a href="user?id=dmor">dmor</a> <a href="item?id=9050332">10 hours ago</a>  | <a href="item?id=9050332">3 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">16.</span></td><td><center><a id="up_9050379" href="vote?for=9050379&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050379"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://www.desmos.com/calculator/9ed6k6wexp">Show HN: A Sierpinski Valentine in 4 equations</a><span class="sitebit comhead"> (desmos.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050379">55 points</span> by <a href="user?id=harmonium1729">harmonium1729</a> <a href="item?id=9050379">9 hours ago</a>  | <a href="item?id=9050379">7 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">17.</span></td><td><center><a id="up_9050601" href="vote?for=9050601&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050601"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://pointersgonewild.com/2015/02/08/my-thesis-a-clearer-picture/">Basic Block Versioning: A Clearer Picture</a><span class="sitebit comhead"> (pointersgonewild.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050601">14 points</span> by <a href="user?id=jwmerrill">jwmerrill</a> <a href="item?id=9050601">8 hours ago</a>  | <a href="item?id=9050601">discuss</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">18.</span></td><td><center><a id="up_9050468" href="vote?for=9050468&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050468"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://hmarco.org/bugs/linux-ASLR-integer-overflow.html">CVE-2015-1593 – Linux ASLR integer overflow: Reducing stack entropy by four</a><span class="sitebit comhead"> (hmarco.org)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050468">16 points</span> by <a href="user?id=adamnemecek">adamnemecek</a> <a href="item?id=9050468">9 hours ago</a>  | <a href="item?id=9050468">2 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">19.</span></td><td><center><a id="up_9049630" href="vote?for=9049630&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049630"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://codecapsule.com/2014/02/12/coding-for-ssds-part-6-a-summary-what-every-programmer-should-know-about-solid-state-drives/">What every programmer should know about solid-state drives</a><span class="sitebit comhead"> (codecapsule.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049630">91 points</span> by <a href="user?id=kozlovsky">kozlovsky</a> <a href="item?id=9049630">14 hours ago</a>  | <a href="item?id=9049630">18 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">20.</span></td><td><center><a id="up_9051591" href="vote?for=9051591&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051591"></span></center></td><td class="title"><span class="deadmark"></span><a href="https://www.owasp.org/index.php/HTML5_Security_Cheat_Sheet" rel="nofollow">HTML5 Security Cheat Sheet – OWASP</a><span class="sitebit comhead"> (owasp.org)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051591">6 points</span> by <a href="user?id=dhruvbhatia">dhruvbhatia</a> <a href="item?id=9051591">1 hour ago</a>  | <a href="item?id=9051591">1 comment</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">21.</span></td><td><center><a id="up_9049597" href="vote?for=9049597&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049597"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://arxiv.org/abs/1412.5567">Deep Speech: Scaling up end-to-end speech recognition</a><span class="sitebit comhead"> (arxiv.org)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049597">53 points</span> by <a href="user?id=lelf">lelf</a> <a href="item?id=9049597">14 hours ago</a>  | <a href="item?id=9049597">4 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">22.</span></td><td><center><a id="up_9049569" href="vote?for=9049569&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049569"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://motherboard.vice.com/read/when-the-future-helps-determine-the-past">When Time Flows Backwards</a><span class="sitebit comhead"> (vice.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049569">46 points</span> by <a href="user?id=DiabloD3">DiabloD3</a> <a href="item?id=9049569">15 hours ago</a>  | <a href="item?id=9049569">8 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">23.</span></td><td><center><a id="up_9049945" href="vote?for=9049945&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049945"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://marc.info/?l=openbsd-tech&amp;m=142356166731390&amp;w=2">Authenticated TLS “contraints” in ntpd(8)</a><span class="sitebit comhead"> (marc.info)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049945">61 points</span> by <a href="user?id=protomyth">protomyth</a> <a href="item?id=9049945">12 hours ago</a>  | <a href="item?id=9049945">5 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">24.</span></td><td><center><a id="up_9050316" href="vote?for=9050316&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050316"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://34cross.in/zero-to-haskell-in-production">Haskell in Production</a><span class="sitebit comhead"> (34cross.in)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050316">50 points</span> by <a href="user?id=arkhamist">arkhamist</a> <a href="item?id=9050316">10 hours ago</a>  | <a href="item?id=9050316">32 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">25.</span></td><td><center><a id="up_9049907" href="vote?for=9049907&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049907"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://technoquarter.blogspot.com/2015/02/openbsd-mail-server.html">OpenBSD Mail Server Intro</a><span class="sitebit comhead"> (technoquarter.blogspot.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049907">54 points</span> by <a href="user?id=protomyth">protomyth</a> <a href="item?id=9049907">12 hours ago</a>  | <a href="item?id=9049907">22 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">26.</span></td><td><center><a id="up_9049348" href="vote?for=9049348&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9049348"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://stupidpythonideas.blogspot.com/2015/02/augmented-assignments-b.html">Understanding Python's augmented assignment (a += b)</a><span class="sitebit comhead"> (stupidpythonideas.blogspot.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9049348">65 points</span> by <a href="user?id=ceronman">ceronman</a> <a href="item?id=9049348">17 hours ago</a>  | <a href="item?id=9049348">13 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">27.</span></td><td><center><a id="up_9051271" href="vote?for=9051271&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9051271"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.bbc.co.uk/newsbeat/31470124">Zane Lowe to join Apple</a><span class="sitebit comhead"> (bbc.co.uk)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9051271">24 points</span> by <a href="user?id=lewispb">lewispb</a> <a href="item?id=9051271">4 hours ago</a>  | <a href="item?id=9051271">6 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">28.</span></td><td><center><a id="up_9050595" href="vote?for=9050595&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050595"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.wsj.com/articles/SB125599860004295449">How a Fight Over a Board Game Monopolized an Economist's Life (2009)</a><span class="sitebit comhead"> (wsj.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050595">8 points</span> by <a href="user?id=ganeumann">ganeumann</a> <a href="item?id=9050595">8 hours ago</a>  | <a href="item?id=9050595">discuss</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">29.</span></td><td><center><a id="up_9050372" href="vote?for=9050372&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050372"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://www.bloomberg.com/graphics/2015-dark-science-of-interrogation/">The Dark Science of Interrogation</a><span class="sitebit comhead"> (bloomberg.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050372">43 points</span> by <a href="user?id=blago">blago</a> <a href="item?id=9050372">9 hours ago</a>  | <a href="item?id=9050372">16 comments</a></td></tr><tr style="height:5px"></tr><tr><td align="right" valign="top" class="title"><span class="rank">30.</span></td><td><center><a id="up_9050009" href="vote?for=9050009&amp;dir=up&amp;goto=news"><div class="votearrow" title="upvote"></div></a><span id="down_9050009"></span></center></td><td class="title"><span class="deadmark"></span><a href="http://docs.topazruby.com/en/latest/">High Performance Ruby in RPython</a><span class="sitebit comhead"> (topazruby.com)</span></td></tr><tr><td colspan="2"></td><td class="subtext"><span class="score" id="score_9050009">35 points</span> by <a href="user?id=renlinx">renlinx</a> <a href="item?id=9050009">11 hours ago</a>  | <a href="item?id=9050009">11 comments</a></td></tr><tr style="height:5px"></tr><tr style="height:10px"></tr><tr><td colspan="2"></td><td class="title"><a href="news?p=2" rel="nofollow">More</a></td></tr></table></td></tr><tr><td><img src="s.gif" height="10" width="0"><table width="100%" cellspacing="0" cellpadding="1"><tr><td bgcolor="#ff6600"></td></tr></table><br>
<center><span class="yclinks"><a href="newsguidelines.html">Guidelines</a> | <a href="newsfaq.html">FAQ</a> | <a href="mailto:hn@ycombinator.com">Support</a> | <a href="lists">Lists</a> | <a href="bookmarklet.html">Bookmarklet</a> | <a href="dmca.html">DMCA</a> | <a href="http://www.ycombinator.com/">Y Combinator</a> | <a href="http://www.ycombinator.com/apply/">Apply</a> | <a href="mailto:hn@ycombinator.com">Contact</a></span><br><br>
<form method="get" action="//hn.algolia.com/">Search: <input type="text" name="q" value="" size="17"></form></center></td></tr></table></center></body></html>`

const COOKIE = "__cfduid=hideyhoneighbor"

func TestRetreivePage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gw := gzip.NewWriter(w)
		defer gw.Close()

		switch r.Method {
		case "HEAD":
			w.Header()["Set-Cookie"] = []string{COOKIE}
			fmt.Fprintln(gw)
			break
		default:
			if cookie, err := r.Cookie("__cfduid"); err == nil {
				if cookie.String() != "__cfduid=hideyhoneighbor" {
					t.Errorf("Cookie not set properly: %s", cookie)
				}
			} else {
				t.Error(err)
			}

			fmt.Fprintf(gw, TestResponse)
			break
		}
	}))

	defer ts.Close()

	c := NewClient(ts.URL)

	p, err := c.RetrievePage("/")

	if err != nil {
		t.Fatal("Error retrieving page", err)
	}

	if len(p.Articles) != 30 {
		t.Fatalf("Not enough articles: %d", len(p.Articles))
	}

	a := p.Articles[0]

	if a.Id != 9050970 {
		t.Errorf("Wrong id for article 0: %d", a.Id)
	}

	if a.Rank != 0 {
		t.Errorf("Wrong rank for article 0: %d", a.Rank)
	}

	if a.Url != "http://websocketd.com/" {
		t.Errorf("Wrong url for article: %s", a.Url)
	}

	if a.Karma != 217 {
		t.Errorf("Wrong karma for article: %d", a.Karma)
	}
}

// vim: set nowrap :
