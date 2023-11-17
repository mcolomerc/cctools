import{_ as n,r as s,o as e,c as o,a as c,b as t,w as i,d as l,e as p}from"./app-0a598c04.js";const r={},d=p(`<h1 id="copy-topics" tabindex="-1"><a class="header-anchor" href="#copy-topics" aria-hidden="true">#</a> Copy Topics</h1><p>Copy Topics metadata to destination cluster.</p><p>Includes Topic configuration and Topic ACLs.</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools copy topics <span class="token parameter variable">--config</span> config.yml
</code></pre></div><p>Usage:</p><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>Command to copy from <span class="token builtin class-name">source</span> Kafka and create destination Topics.

Usage:
  cctools copy topics <span class="token punctuation">[</span>flags<span class="token punctuation">]</span>

Aliases:
  topics, topic-cp, tpic-cp, tpc

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> topics

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>  
</code></pre></div><h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><ul><li>Using Topic copyer Configuration to exclude some topics.</li></ul><p>All topics names containing <code>_confluent</code> will be excluded.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">copy</span><span class="token punctuation">:</span> 
  <span class="token key atrule">topics</span><span class="token punctuation">:</span>
    <span class="token key atrule">exclude</span><span class="token punctuation">:</span> _confluent 
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><ul><li>Topic ACLs - Principals Mapping</li></ul><p>All the Topic ACLs where <code>principal: User:test</code> will be created as <code>principal: User:sa-xyroox</code> on the Destination.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">principals</span><span class="token punctuation">:</span>
  <span class="token punctuation">-</span> <span class="token key atrule">&quot;test&quot;</span><span class="token punctuation">:</span> <span class="token string">&quot;sa-xyroox&quot;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div>`,13);function u(m,g){const a=s("RouterLink");return e(),o("div",null,[d,c("p",null,[t(a,{to:"/commands/config/"},{default:i(()=>[l("Configuration")]),_:1})])])}const h=n(r,[["render",u],["__file","topics.html.vue"]]);export{h as default};
