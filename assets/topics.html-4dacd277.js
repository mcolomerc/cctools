import{_ as c,r as p,o as i,c as l,a as s,d as n,b as t,w as e,e as o}from"./app-0a598c04.js";const r={},u=o(`<h1 id="import-topics" tabindex="-1"><a class="header-anchor" href="#import-topics" aria-hidden="true">#</a> Import Topics</h1><p>Import Topics metadata from source files and create destination Topics into Destination cluster</p><p>Includes Topic configuration and Topic ACLs.</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token function">import</span> topics --help\`
</code></pre></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>Command to <span class="token function">import</span> from <span class="token builtin class-name">source</span> files and create destination Topics.

Usage:
  cctools <span class="token function">import</span> topics <span class="token punctuation">[</span>flags<span class="token punctuation">]</span> 

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> topics

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>
</code></pre></div>`,5),d={class:"custom-container tip"},k=s("p",{class:"custom-container-title"},"TIP",-1),m=o(`<h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><p>Destination cluster.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">destination</span><span class="token punctuation">:</span> 
  <span class="token key atrule">kafka</span><span class="token punctuation">:</span>
    <span class="token key atrule">bootstrapServer</span><span class="token punctuation">:</span> &lt;bootstrap_server<span class="token punctuation">&gt;</span>.confluent.cloud<span class="token punctuation">:</span><span class="token number">9092</span>
    <span class="token key atrule">clientProps</span><span class="token punctuation">:</span>
      <span class="token punctuation">-</span> <span class="token key atrule">sasl.mechanisms</span><span class="token punctuation">:</span> PLAIN
      <span class="token punctuation">-</span> <span class="token key atrule">security.protocol</span><span class="token punctuation">:</span> SASL_SSL
      <span class="token punctuation">-</span> <span class="token key atrule">sasl.username</span><span class="token punctuation">:</span> &lt;API_KEY<span class="token punctuation">&gt;</span>
      <span class="token punctuation">-</span> <span class="token key atrule">sasl.password</span><span class="token punctuation">:</span> &lt;API_SECRET<span class="token punctuation">&gt;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,3);function f(h,v){const a=p("RouterLink");return i(),l("div",null,[u,s("div",d,[k,s("p",null,[n("Works with exported JSON files. See "),t(a,{to:"/commands/export/topics.html"},{default:e(()=>[n("Export Topics")]),_:1}),n(" for more information.")])]),m,s("p",null,[n("See "),t(a,{to:"/commands/config/"},{default:e(()=>[n("Configuration")]),_:1})])])}const _=c(r,[["render",f],["__file","topics.html.vue"]]);export{_ as default};
