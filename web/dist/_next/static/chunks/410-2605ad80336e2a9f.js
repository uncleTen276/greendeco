"use strict";(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[410],{622:function(t,e,n){/**
 * @license React
 * react-jsx-runtime.production.min.js
 *
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var i=n(2265),r=Symbol.for("react.element"),o=Symbol.for("react.fragment"),s=Object.prototype.hasOwnProperty,u=i.__SECRET_INTERNALS_DO_NOT_USE_OR_YOU_WILL_BE_FIRED.ReactCurrentOwner,a={key:!0,ref:!0,__self:!0,__source:!0};function l(t,e,n){var i,o={},l=null,c=null;for(i in void 0!==n&&(l=""+n),void 0!==e.key&&(l=""+e.key),void 0!==e.ref&&(c=e.ref),e)s.call(e,i)&&!a.hasOwnProperty(i)&&(o[i]=e[i]);if(t&&t.defaultProps)for(i in e=t.defaultProps)void 0===o[i]&&(o[i]=e[i]);return{$$typeof:r,type:t,key:l,ref:c,props:o,_owner:u.current}}e.Fragment=o,e.jsx=l,e.jsxs=l},7437:function(t,e,n){t.exports=n(622)},8202:function(t,e,n){n.d(e,{j:function(){return s}});var i=n(9492),r=n(6504);class o extends i.l{constructor(){super(),this.setup=t=>{if(!r.sk&&window.addEventListener){let e=()=>t();return window.addEventListener("visibilitychange",e,!1),window.addEventListener("focus",e,!1),()=>{window.removeEventListener("visibilitychange",e),window.removeEventListener("focus",e)}}}}onSubscribe(){this.cleanup||this.setEventListener(this.setup)}onUnsubscribe(){if(!this.hasListeners()){var t;null==(t=this.cleanup)||t.call(this),this.cleanup=void 0}}setEventListener(t){var e;this.setup=t,null==(e=this.cleanup)||e.call(this),this.cleanup=t(t=>{"boolean"==typeof t?this.setFocused(t):this.onFocus()})}setFocused(t){let e=this.focused!==t;e&&(this.focused=t,this.onFocus())}onFocus(){this.listeners.forEach(({listener:t})=>{t()})}isFocused(){return"boolean"==typeof this.focused?this.focused:"undefined"==typeof document||[void 0,"visible","prerender"].includes(document.visibilityState)}}let s=new o},8810:function(t,e,n){n.d(e,{_:function(){return i}});let i=console},172:function(t,e,n){n.d(e,{R:function(){return a},m:function(){return u}});var i=n(8810),r=n(7156),o=n(1909),s=n(3238);class u extends o.F{constructor(t){super(),this.defaultOptions=t.defaultOptions,this.mutationId=t.mutationId,this.mutationCache=t.mutationCache,this.logger=t.logger||i._,this.observers=[],this.state=t.state||a(),this.setOptions(t.options),this.scheduleGc()}setOptions(t){this.options={...this.defaultOptions,...t},this.updateCacheTime(this.options.cacheTime)}get meta(){return this.options.meta}setState(t){this.dispatch({type:"setState",state:t})}addObserver(t){this.observers.includes(t)||(this.observers.push(t),this.clearGcTimeout(),this.mutationCache.notify({type:"observerAdded",mutation:this,observer:t}))}removeObserver(t){this.observers=this.observers.filter(e=>e!==t),this.scheduleGc(),this.mutationCache.notify({type:"observerRemoved",mutation:this,observer:t})}optionalRemove(){this.observers.length||("loading"===this.state.status?this.scheduleGc():this.mutationCache.remove(this))}continue(){var t,e;return null!=(t=null==(e=this.retryer)?void 0:e.continue())?t:this.execute()}async execute(){var t,e,n,i,r,o,u,a,l,c,h,f,d,p,v,y,b,m,w,C;let g="loading"===this.state.status;try{if(!g){this.dispatch({type:"loading",variables:this.options.variables}),await (null==(l=(c=this.mutationCache.config).onMutate)?void 0:l.call(c,this.state.variables,this));let t=await (null==(h=(f=this.options).onMutate)?void 0:h.call(f,this.state.variables));t!==this.state.context&&this.dispatch({type:"loading",context:t,variables:this.state.variables})}let d=await (()=>{var t;return this.retryer=(0,s.Mz)({fn:()=>this.options.mutationFn?this.options.mutationFn(this.state.variables):Promise.reject("No mutationFn found"),onFail:(t,e)=>{this.dispatch({type:"failed",failureCount:t,error:e})},onPause:()=>{this.dispatch({type:"pause"})},onContinue:()=>{this.dispatch({type:"continue"})},retry:null!=(t=this.options.retry)?t:0,retryDelay:this.options.retryDelay,networkMode:this.options.networkMode}),this.retryer.promise})();return await (null==(t=(e=this.mutationCache.config).onSuccess)?void 0:t.call(e,d,this.state.variables,this.state.context,this)),await (null==(n=(i=this.options).onSuccess)?void 0:n.call(i,d,this.state.variables,this.state.context)),await (null==(r=(o=this.mutationCache.config).onSettled)?void 0:r.call(o,d,null,this.state.variables,this.state.context,this)),await (null==(u=(a=this.options).onSettled)?void 0:u.call(a,d,null,this.state.variables,this.state.context)),this.dispatch({type:"success",data:d}),d}catch(t){try{throw await (null==(d=(p=this.mutationCache.config).onError)?void 0:d.call(p,t,this.state.variables,this.state.context,this)),await (null==(v=(y=this.options).onError)?void 0:v.call(y,t,this.state.variables,this.state.context)),await (null==(b=(m=this.mutationCache.config).onSettled)?void 0:b.call(m,void 0,t,this.state.variables,this.state.context,this)),await (null==(w=(C=this.options).onSettled)?void 0:w.call(C,void 0,t,this.state.variables,this.state.context)),t}finally{this.dispatch({type:"error",error:t})}}}dispatch(t){this.state=(e=>{switch(t.type){case"failed":return{...e,failureCount:t.failureCount,failureReason:t.error};case"pause":return{...e,isPaused:!0};case"continue":return{...e,isPaused:!1};case"loading":return{...e,context:t.context,data:void 0,failureCount:0,failureReason:null,error:null,isPaused:!(0,s.Kw)(this.options.networkMode),status:"loading",variables:t.variables};case"success":return{...e,data:t.data,failureCount:0,failureReason:null,error:null,status:"success",isPaused:!1};case"error":return{...e,data:void 0,error:t.error,failureCount:e.failureCount+1,failureReason:t.error,isPaused:!1,status:"error"};case"setState":return{...e,...t.state}}})(this.state),r.V.batch(()=>{this.observers.forEach(e=>{e.onMutationUpdate(t)}),this.mutationCache.notify({mutation:this,type:"updated",action:t})})}}function a(){return{context:void 0,data:void 0,error:null,failureCount:0,failureReason:null,isPaused:!1,status:"idle",variables:void 0}}},7156:function(t,e,n){n.d(e,{V:function(){return r}});var i=n(6504);let r=function(){let t=[],e=0,n=t=>{t()},r=t=>{t()},o=r=>{e?t.push(r):(0,i.A4)(()=>{n(r)})},s=()=>{let e=t;t=[],e.length&&(0,i.A4)(()=>{r(()=>{e.forEach(t=>{n(t)})})})};return{batch:t=>{let n;e++;try{n=t()}finally{--e||s()}return n},batchCalls:t=>(...e)=>{o(()=>{t(...e)})},schedule:o,setNotifyFunction:t=>{n=t},setBatchNotifyFunction:t=>{r=t}}}()},3864:function(t,e,n){n.d(e,{N:function(){return u}});var i=n(9492),r=n(6504);let o=["online","offline"];class s extends i.l{constructor(){super(),this.setup=t=>{if(!r.sk&&window.addEventListener){let e=()=>t();return o.forEach(t=>{window.addEventListener(t,e,!1)}),()=>{o.forEach(t=>{window.removeEventListener(t,e)})}}}}onSubscribe(){this.cleanup||this.setEventListener(this.setup)}onUnsubscribe(){if(!this.hasListeners()){var t;null==(t=this.cleanup)||t.call(this),this.cleanup=void 0}}setEventListener(t){var e;this.setup=t,null==(e=this.cleanup)||e.call(this),this.cleanup=t(t=>{"boolean"==typeof t?this.setOnline(t):this.onOnline()})}setOnline(t){let e=this.online!==t;e&&(this.online=t,this.onOnline())}onOnline(){this.listeners.forEach(({listener:t})=>{t()})}isOnline(){return"boolean"==typeof this.online?this.online:"undefined"==typeof navigator||void 0===navigator.onLine||navigator.onLine}}let u=new s},1909:function(t,e,n){n.d(e,{F:function(){return r}});var i=n(6504);class r{destroy(){this.clearGcTimeout()}scheduleGc(){this.clearGcTimeout(),(0,i.PN)(this.cacheTime)&&(this.gcTimeout=setTimeout(()=>{this.optionalRemove()},this.cacheTime))}updateCacheTime(t){this.cacheTime=Math.max(this.cacheTime||0,null!=t?t:i.sk?1/0:3e5)}clearGcTimeout(){this.gcTimeout&&(clearTimeout(this.gcTimeout),this.gcTimeout=void 0)}}},3238:function(t,e,n){n.d(e,{DV:function(){return l},Kw:function(){return u},Mz:function(){return c}});var i=n(8202),r=n(3864),o=n(6504);function s(t){return Math.min(1e3*2**t,3e4)}function u(t){return(null!=t?t:"online")!=="online"||r.N.isOnline()}class a{constructor(t){this.revert=null==t?void 0:t.revert,this.silent=null==t?void 0:t.silent}}function l(t){return t instanceof a}function c(t){let e,n,l,c=!1,h=0,f=!1,d=new Promise((t,e)=>{n=t,l=e}),p=()=>!i.j.isFocused()||"always"!==t.networkMode&&!r.N.isOnline(),v=i=>{f||(f=!0,null==t.onSuccess||t.onSuccess(i),null==e||e(),n(i))},y=n=>{f||(f=!0,null==t.onError||t.onError(n),null==e||e(),l(n))},b=()=>new Promise(n=>{e=t=>{let e=f||!p();return e&&n(t),e},null==t.onPause||t.onPause()}).then(()=>{e=void 0,f||null==t.onContinue||t.onContinue()}),m=()=>{let e;if(!f){try{e=t.fn()}catch(t){e=Promise.reject(t)}Promise.resolve(e).then(v).catch(e=>{var n,i;if(f)return;let r=null!=(n=t.retry)?n:3,u=null!=(i=t.retryDelay)?i:s,a="function"==typeof u?u(h,e):u,l=!0===r||"number"==typeof r&&h<r||"function"==typeof r&&r(h,e);if(c||!l){y(e);return}h++,null==t.onFail||t.onFail(h,e),(0,o.Gh)(a).then(()=>{if(p())return b()}).then(()=>{c?y(e):m()})})}};return u(t.networkMode)?m():b().then(m),{promise:d,cancel:e=>{f||(y(new a(e)),null==t.abort||t.abort())},continue:()=>{let t=null==e?void 0:e();return t?d:Promise.resolve()},cancelRetry:()=>{c=!0},continueRetry:()=>{c=!1}}}},9492:function(t,e,n){n.d(e,{l:function(){return i}});class i{constructor(){this.listeners=new Set,this.subscribe=this.subscribe.bind(this)}subscribe(t){let e={listener:t};return this.listeners.add(e),this.onSubscribe(),()=>{this.listeners.delete(e),this.onUnsubscribe()}}hasListeners(){return this.listeners.size>0}onSubscribe(){}onUnsubscribe(){}}},6504:function(t,e,n){n.d(e,{A4:function(){return O},G9:function(){return S},Gh:function(){return E},I6:function(){return c},Kp:function(){return u},PN:function(){return s},Rm:function(){return d},SE:function(){return o},VS:function(){return b},X7:function(){return f},ZT:function(){return r},_v:function(){return a},_x:function(){return h},lV:function(){return l},oE:function(){return x},sk:function(){return i},to:function(){return v},yF:function(){return p}});let i="undefined"==typeof window||"Deno"in window;function r(){}function o(t,e){return"function"==typeof t?t(e):t}function s(t){return"number"==typeof t&&t>=0&&t!==1/0}function u(t,e){return Math.max(t+(e||0)-Date.now(),0)}function a(t,e,n){return g(t)?"function"==typeof e?{...n,queryKey:t,queryFn:e}:{...e,queryKey:t}:t}function l(t,e,n){return g(t)?"function"==typeof e?{...n,mutationKey:t,mutationFn:e}:{...e,mutationKey:t}:"function"==typeof t?{...e,mutationFn:t}:{...t}}function c(t,e,n){return g(t)?[{...e,queryKey:t},n]:[t||{},e]}function h(t,e){let{type:n="all",exact:i,fetchStatus:r,predicate:o,queryKey:s,stale:u}=t;if(g(s)){if(i){if(e.queryHash!==d(s,e.options))return!1}else{if(!y(e.queryKey,s))return!1}}if("all"!==n){let t=e.isActive();if("active"===n&&!t||"inactive"===n&&t)return!1}return("boolean"!=typeof u||e.isStale()===u)&&(void 0===r||r===e.state.fetchStatus)&&(!o||!!o(e))}function f(t,e){let{exact:n,fetching:i,predicate:r,mutationKey:o}=t;if(g(o)){if(!e.options.mutationKey)return!1;if(n){if(p(e.options.mutationKey)!==p(o))return!1}else{if(!y(e.options.mutationKey,o))return!1}}return("boolean"!=typeof i||"loading"===e.state.status===i)&&(!r||!!r(e))}function d(t,e){let n=(null==e?void 0:e.queryKeyHashFn)||p;return n(t)}function p(t){return JSON.stringify(t,(t,e)=>w(e)?Object.keys(e).sort().reduce((t,n)=>(t[n]=e[n],t),{}):e)}function v(t,e){return y(t,e)}function y(t,e){return t===e||typeof t==typeof e&&!!t&&!!e&&"object"==typeof t&&"object"==typeof e&&!Object.keys(e).some(n=>!y(t[n],e[n]))}function b(t,e){if(t&&!e||e&&!t)return!1;for(let n in t)if(t[n]!==e[n])return!1;return!0}function m(t){return Array.isArray(t)&&t.length===Object.keys(t).length}function w(t){if(!C(t))return!1;let e=t.constructor;if(void 0===e)return!0;let n=e.prototype;return!!(C(n)&&n.hasOwnProperty("isPrototypeOf"))}function C(t){return"[object Object]"===Object.prototype.toString.call(t)}function g(t){return Array.isArray(t)}function E(t){return new Promise(e=>{setTimeout(e,t)})}function O(t){E(0).then(t)}function S(){if("function"==typeof AbortController)return new AbortController}function x(t,e,n){return null!=n.isDataEqual&&n.isDataEqual(t,e)?t:"function"==typeof n.structuralSharing?n.structuralSharing(t,e):!1!==n.structuralSharing?function t(e,n){if(e===n)return e;let i=m(e)&&m(n);if(i||w(e)&&w(n)){let r=i?e.length:Object.keys(e).length,o=i?n:Object.keys(n),s=o.length,u=i?[]:{},a=0;for(let r=0;r<s;r++){let s=i?r:o[r];u[s]=t(e[s],n[s]),u[s]===e[s]&&a++}return r===s&&a===r?e:u}return n}(t,e):e}},165:function(t,e,n){n.d(e,{NL:function(){return u},aH:function(){return a}});var i=n(2265);let r=i.createContext(void 0),o=i.createContext(!1);function s(t,e){return t||(e&&"undefined"!=typeof window?(window.ReactQueryClientContext||(window.ReactQueryClientContext=r),window.ReactQueryClientContext):r)}let u=({context:t}={})=>{let e=i.useContext(s(t,i.useContext(o)));if(!e)throw Error("No QueryClient set, use QueryClientProvider to set one");return e},a=({client:t,children:e,context:n,contextSharing:r=!1})=>{i.useEffect(()=>(t.mount(),()=>{t.unmount()}),[t]);let u=s(n,r);return i.createElement(o.Provider,{value:!n&&r},i.createElement(u.Provider,{value:t},e))}}}]);