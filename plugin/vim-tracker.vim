" Initialize the channel
if !exists('s:trackerVimJobId')
	let s:trackerVimJobId = 0
endif

" The path to the binary that was created out of 'cargo build' or 'cargo build --release". This will generally be 'target/release/name'
let s:bin = 'trackervim'

" Entry point. Initialize RPC. If it succeeds, then attach commands to the `rpcnotify` invocations.
function! s:connect()
  let id = s:initRpc()

  if 0 == id
    echoerr "tracker-vim: cannot start rpc process"
  elseif -1 == id
    echoerr "tracker-vim: rpc process is not executable"
  else
    " Mutate our jobId variable to hold the channel ID
    let s:trackerVimJobId = id

    // TODO: Configure commands to their RPC invocations.
  endif
endfunction

" Initialize RPC
function! s:initRpc()
  if s:trackerVimJobId == 0
    let jobid = jobstart([s:bin], { 'rpc': v:true })
    return jobid
  else
    return s:trackerVimJobId
  endif
endfunction
