import{_ as o,r as l,o as c,c as r,a as e,b as n,w as t,d as a,e as p}from"./app-0a598c04.js";const i={},u=p(`<h1 id="export" tabindex="-1"><a class="header-anchor" href="#export" aria-hidden="true">#</a> Export</h1><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token builtin class-name">export</span> <span class="token parameter variable">--help</span>
</code></pre></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>Command to <span class="token builtin class-name">export</span> cluster information.

Usage:
  cctools <span class="token builtin class-name">export</span> <span class="token punctuation">[</span>flags<span class="token punctuation">]</span>
  cctools <span class="token builtin class-name">export</span> <span class="token punctuation">[</span>command<span class="token punctuation">]</span> 
  
Available Commands:
  schemas     Export Schemas Info
  topics      Export Topics Info

Flags:
  -h, <span class="token parameter variable">--help</span>            <span class="token builtin class-name">help</span> <span class="token keyword">for</span> <span class="token builtin class-name">export</span>
  -o, <span class="token parameter variable">--output</span> string   Output format. Possible values: json, yaml, hcl, cfk, clink

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span> 

Use <span class="token string">&quot;cctools export [command] --help&quot;</span> <span class="token keyword">for</span> <span class="token function">more</span> information about a command.
</code></pre></div><h2 id="output-folder" tabindex="-1"><a class="header-anchor" href="#output-folder" aria-hidden="true">#</a> Output folder</h2><p>Configure the output folder, it will be created if it does not exist.</p><p>Example: All the export files will be stored into the <code>output</code> folder (it will be created if necessary).</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">export</span><span class="token punctuation">:</span> 
  <span class="token key atrule">output</span><span class="token punctuation">:</span> output 
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><ol><li><p>Each <code>resource</code> will create a folder inside the <code>output</code> target.</p></li><li><p>Each exporter will create a folder inside the <code>resource</code> folder.</p></li></ol><p><strong>Example</strong>: Exporting Topics to JSON will generate: <code>output/topics/json/topics.json</code></p><h2 id="resources" tabindex="-1"><a class="header-anchor" href="#resources" aria-hidden="true">#</a> Resources</h2>`,10),d=e("h2",{id:"exporters",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#exporters","aria-hidden":"true"},"#"),a(" Exporters")],-1),m=e("p",null,[e("code",null,"--output"),a(" flag is required.")],-1);function h(f,x){const s=l("RouterLink");return c(),r("div",null,[u,e("ul",null,[e("li",null,[e("p",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("Topics")]),_:1})])]),e("li",null,[e("p",null,[n(s,{to:"/commands/export/consumer-groups.html"},{default:t(()=>[a("Consumer Groups")]),_:1})])]),e("li",null,[e("p",null,[n(s,{to:"/commands/export/schemas.html"},{default:t(()=>[a("Schemas")]),_:1})])])]),d,m,e("ul",null,[e("li",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("JSON")]),_:1})]),e("li",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("YAML")]),_:1})]),e("li",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("HCL")]),_:1})]),e("li",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("Confluent Cloud")]),_:1})]),e("li",null,[n(s,{to:"/commands/export/topics.html"},{default:t(()=>[a("Confluent Cloud Link")]),_:1})])])])}const b=o(i,[["render",h],["__file","export.html.vue"]]);export{b as default};
