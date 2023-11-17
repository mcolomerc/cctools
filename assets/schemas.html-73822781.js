import{_ as c,r as i,o as p,c as l,a as n,d as s,b as e,w as t,e as o}from"./app-0a598c04.js";const r={},u=o(`<h1 id="import-schemas" tabindex="-1"><a class="header-anchor" href="#import-schemas" aria-hidden="true">#</a> Import Schemas</h1><p>Import Schemas from source files and create destination Subjects into Destination Schema Registry.</p><p>Subjects are created on the destination keeping the schema <code>id</code> and the schema <code>version</code>.</p><div class="custom-container warning"><p class="custom-container-title">WARNING</p><p>If the schema <code>id</code> already exists on the destination, Schema Registry call will fail with an error when switching to <code>IMPORT</code>.</p></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>cctools <span class="token function">import</span> schemas --help\`
</code></pre></div><div class="language-bash" data-ext="sh"><pre class="language-bash"><code>Command to <span class="token function">import</span> from <span class="token builtin class-name">source</span> files and create destination Schemas.

Usage:
  cctools <span class="token function">import</span> schemas <span class="token punctuation">[</span>flags<span class="token punctuation">]</span>

Flags:
  -h, <span class="token parameter variable">--help</span>   <span class="token builtin class-name">help</span> <span class="token keyword">for</span> schemas

Global Flags:
  -c, <span class="token parameter variable">--config</span> string   config <span class="token function">file</span>
</code></pre></div>`,6),d={class:"custom-container tip"},m=n("p",{class:"custom-container-title"},"TIP",-1),h=o(`<p>Example input JSON file:</p><div class="language-json line-numbers-mode" data-ext="json"><pre class="language-json"><code><span class="token punctuation">{</span>
    <span class="token property">&quot;schema&quot;</span><span class="token operator">:</span><span class="token string">&quot;{\\&quot;type\\&quot;:\\&quot;record\\&quot;,\\&quot;name\\&quot;:\\&quot;value_a1\\&quot;,\\&quot;namespace\\&quot;:\\&quot;com.mycorp.mynamespace\\&quot;,\\&quot;fields\\&quot;:[{\\&quot;name\\&quot;:\\&quot;field1\\&quot;,\\&quot;type\\&quot;:\\&quot;string\\&quot;}]}&quot;</span><span class="token punctuation">,</span>
    <span class="token property">&quot;version&quot;</span><span class="token operator">:</span><span class="token number">19</span><span class="token punctuation">,</span>
    <span class="token property">&quot;id&quot;</span><span class="token operator">:</span><span class="token number">101010</span><span class="token punctuation">,</span>
    <span class="token property">&quot;schemaType&quot;</span><span class="token operator">:</span> <span class="token string">&quot;AVRO&quot;</span><span class="token punctuation">,</span>
    <span class="token property">&quot;subject&quot;</span><span class="token operator">:</span> <span class="token string">&quot;my-value&quot;</span> 
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="configuration" tabindex="-1"><a class="header-anchor" href="#configuration" aria-hidden="true">#</a> Configuration</h2><p>Destination cluster.</p><div class="language-yaml line-numbers-mode" data-ext="yml"><pre class="language-yaml"><code><span class="token key atrule">destination</span><span class="token punctuation">:</span> 
  <span class="token key atrule">schemaRegistry</span><span class="token punctuation">:</span>
    <span class="token key atrule">endpointUrl</span><span class="token punctuation">:</span> &lt;SCHEMA_REGISTRY_URL<span class="token punctuation">&gt;</span>
    <span class="token key atrule">credentials</span><span class="token punctuation">:</span>
      <span class="token key atrule">key</span><span class="token punctuation">:</span> &lt;USER<span class="token punctuation">&gt;</span>
      <span class="token key atrule">secret</span><span class="token punctuation">:</span> &lt;PASSWORD<span class="token punctuation">&gt;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5);function k(v,g){const a=i("RouterLink");return p(),l("div",null,[u,n("div",d,[m,n("p",null,[s("Works with JSON files. See "),e(a,{to:"/commands/export/schemas.html"},{default:t(()=>[s("Export Schemas")]),_:1}),s(" for more information.")])]),h,n("p",null,[s("See "),e(a,{to:"/commands/config/"},{default:t(()=>[s("Configuration")]),_:1})])])}const f=c(r,[["render",k],["__file","schemas.html.vue"]]);export{f as default};
