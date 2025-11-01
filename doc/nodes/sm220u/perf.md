  branch-instructions OR branches                    [Hardware event]
  branch-misses                                      [Hardware event]
  bus-cycles                                         [Hardware event]
  cache-misses                                       [Hardware event]
  cache-references                                   [Hardware event]
  cpu-cycles OR cycles                               [Hardware event]
  instructions                                       [Hardware event]
  ref-cycles                                         [Hardware event]
  alignment-faults                                   [Software event]
  bpf-output                                         [Software event]
  cgroup-switches                                    [Software event]
  context-switches OR cs                             [Software event]
  cpu-clock                                          [Software event]
  cpu-migrations OR migrations                       [Software event]
  dummy                                              [Software event]
  emulation-faults                                   [Software event]
  major-faults                                       [Software event]
  minor-faults                                       [Software event]
  page-faults OR faults                              [Software event]
  task-clock                                         [Software event]
  duration_time                                      [Tool event]
  L1-dcache-load-misses                              [Hardware cache event]
  L1-dcache-loads                                    [Hardware cache event]
  L1-dcache-stores                                   [Hardware cache event]
  L1-icache-load-misses                              [Hardware cache event]
  LLC-load-misses                                    [Hardware cache event]
  LLC-loads                                          [Hardware cache event]
  LLC-store-misses                                   [Hardware cache event]
  LLC-stores                                         [Hardware cache event]
  branch-load-misses                                 [Hardware cache event]
  branch-loads                                       [Hardware cache event]
  dTLB-load-misses                                   [Hardware cache event]
  dTLB-loads                                         [Hardware cache event]
  dTLB-store-misses                                  [Hardware cache event]
  dTLB-stores                                        [Hardware cache event]
  iTLB-load-misses                                   [Hardware cache event]
  node-load-misses                                   [Hardware cache event]
  node-loads                                         [Hardware cache event]
  node-store-misses                                  [Hardware cache event]
  node-stores                                        [Hardware cache event]
  branch-instructions OR cpu/branch-instructions/    [Kernel PMU event]
  branch-misses OR cpu/branch-misses/                [Kernel PMU event]
  bus-cycles OR cpu/bus-cycles/                      [Kernel PMU event]
  cache-misses OR cpu/cache-misses/                  [Kernel PMU event]
  cache-references OR cpu/cache-references/          [Kernel PMU event]
  cpu-cycles OR cpu/cpu-cycles/                      [Kernel PMU event]
  instructions OR cpu/instructions/                  [Kernel PMU event]
  mem-loads OR cpu/mem-loads/                        [Kernel PMU event]
  mem-stores OR cpu/mem-stores/                      [Kernel PMU event]
  ref-cycles OR cpu/ref-cycles/                      [Kernel PMU event]
  slots OR cpu/slots/                                [Kernel PMU event]
  topdown-bad-spec OR cpu/topdown-bad-spec/          [Kernel PMU event]
  topdown-be-bound OR cpu/topdown-be-bound/          [Kernel PMU event]
  topdown-fe-bound OR cpu/topdown-fe-bound/          [Kernel PMU event]
  topdown-retiring OR cpu/topdown-retiring/          [Kernel PMU event]
  cstate_core/c1-residency/                          [Kernel PMU event]
  cstate_core/c6-residency/                          [Kernel PMU event]
  cstate_pkg/c2-residency/                           [Kernel PMU event]
  cstate_pkg/c6-residency/                           [Kernel PMU event]
  intel_bts//                                        [Kernel PMU event]
  intel_pt//                                         [Kernel PMU event]
  msr/aperf/                                         [Kernel PMU event]
  msr/cpu_thermal_margin/                            [Kernel PMU event]
  msr/mperf/                                         [Kernel PMU event]
  msr/pperf/                                         [Kernel PMU event]
  msr/smi/                                           [Kernel PMU event]
  msr/tsc/                                           [Kernel PMU event]
  power/energy-pkg/                                  [Kernel PMU event]
  power/energy-ram/                                  [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_0/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_0/ioclk/                   [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_1/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_1/ioclk/                   [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_2/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_2/ioclk/                   [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_3/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_3/ioclk/                   [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_4/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_4/ioclk/                   [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port0/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port1/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port2/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port3/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port4/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port5/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port6/             [Kernel PMU event]
  uncore_iio_free_running_5/bw_in_port7/             [Kernel PMU event]
  uncore_iio_free_running_5/ioclk/                   [Kernel PMU event]
  uncore_imc_0/cas_count_read/                       [Kernel PMU event]
  uncore_imc_0/cas_count_write/                      [Kernel PMU event]
  uncore_imc_0/clockticks/                           [Kernel PMU event]
  uncore_imc_1/cas_count_read/                       [Kernel PMU event]
  uncore_imc_1/cas_count_write/                      [Kernel PMU event]
  uncore_imc_1/clockticks/                           [Kernel PMU event]
  uncore_imc_10/cas_count_read/                      [Kernel PMU event]
  uncore_imc_10/cas_count_write/                     [Kernel PMU event]
  uncore_imc_10/clockticks/                          [Kernel PMU event]
  uncore_imc_11/cas_count_read/                      [Kernel PMU event]
  uncore_imc_11/cas_count_write/                     [Kernel PMU event]
  uncore_imc_11/clockticks/                          [Kernel PMU event]
  uncore_imc_2/cas_count_read/                       [Kernel PMU event]
  uncore_imc_2/cas_count_write/                      [Kernel PMU event]
  uncore_imc_2/clockticks/                           [Kernel PMU event]
  uncore_imc_3/cas_count_read/                       [Kernel PMU event]
  uncore_imc_3/cas_count_write/                      [Kernel PMU event]
  uncore_imc_3/clockticks/                           [Kernel PMU event]
  uncore_imc_4/cas_count_read/                       [Kernel PMU event]
  uncore_imc_4/cas_count_write/                      [Kernel PMU event]
  uncore_imc_4/clockticks/                           [Kernel PMU event]
  uncore_imc_5/cas_count_read/                       [Kernel PMU event]
  uncore_imc_5/cas_count_write/                      [Kernel PMU event]
  uncore_imc_5/clockticks/                           [Kernel PMU event]
  uncore_imc_6/cas_count_read/                       [Kernel PMU event]
  uncore_imc_6/cas_count_write/                      [Kernel PMU event]
  uncore_imc_6/clockticks/                           [Kernel PMU event]
  uncore_imc_7/cas_count_read/                       [Kernel PMU event]
  uncore_imc_7/cas_count_write/                      [Kernel PMU event]
  uncore_imc_7/clockticks/                           [Kernel PMU event]
  uncore_imc_8/cas_count_read/                       [Kernel PMU event]
  uncore_imc_8/cas_count_write/                      [Kernel PMU event]
  uncore_imc_8/clockticks/                           [Kernel PMU event]
  uncore_imc_9/cas_count_read/                       [Kernel PMU event]
  uncore_imc_9/cas_count_write/                      [Kernel PMU event]
  uncore_imc_9/clockticks/                           [Kernel PMU event]
  uncore_imc_free_running_0/dclk/                    [Kernel PMU event]
  uncore_imc_free_running_0/ddrt_read/               [Kernel PMU event]
  uncore_imc_free_running_0/ddrt_write/              [Kernel PMU event]
  uncore_imc_free_running_0/read/                    [Kernel PMU event]
  uncore_imc_free_running_0/write/                   [Kernel PMU event]
  uncore_imc_free_running_1/dclk/                    [Kernel PMU event]
  uncore_imc_free_running_1/ddrt_read/               [Kernel PMU event]
  uncore_imc_free_running_1/ddrt_write/              [Kernel PMU event]
  uncore_imc_free_running_1/read/                    [Kernel PMU event]
  uncore_imc_free_running_1/write/                   [Kernel PMU event]
  uncore_imc_free_running_2/dclk/                    [Kernel PMU event]
  uncore_imc_free_running_2/ddrt_read/               [Kernel PMU event]
  uncore_imc_free_running_2/ddrt_write/              [Kernel PMU event]
  uncore_imc_free_running_2/read/                    [Kernel PMU event]
  uncore_imc_free_running_2/write/                   [Kernel PMU event]
  uncore_imc_free_running_3/dclk/                    [Kernel PMU event]
  uncore_imc_free_running_3/ddrt_read/               [Kernel PMU event]
  uncore_imc_free_running_3/ddrt_write/              [Kernel PMU event]
  uncore_imc_free_running_3/read/                    [Kernel PMU event]
  uncore_imc_free_running_3/write/                   [Kernel PMU event]

cache:
  l1d.replacement                                   
       [Counts the number of cache lines replaced in L1 data cache]
  l1d_pend_miss.fb_full                             
       [Number of cycles a demand request has waited due to L1D Fill Buffer
        (FB) unavailability]
  l1d_pend_miss.fb_full_periods                     
       [Number of phases a demand request has waited due to L1D Fill Buffer
        (FB) unavailablability]
  l1d_pend_miss.l2_stall                            
       [Number of cycles a demand request has waited due to L1D due to lack of
        L2 resources]
  l1d_pend_miss.pending                             
       [Number of L1D misses that are outstanding]
  l1d_pend_miss.pending_cycles                      
       [Cycles with L1D load Misses outstanding]
  l2_lines_in.all                                   
       [L2 cache lines filling L2]
  l2_lines_out.non_silent                           
       [Cache lines that are evicted by L2 cache when triggered by an L2 cache
        fill]
  l2_lines_out.silent                               
       [Non-modified cache lines that are silently dropped by L2 cache when
        triggered by an L2 cache fill]
  l2_rqsts.all_code_rd                              
       [L2 code requests]
  l2_rqsts.all_demand_data_rd                       
       [Demand Data Read requests]
  l2_rqsts.all_demand_miss                          
       [Demand requests that miss L2 cache]
  l2_rqsts.all_rfo                                  
       [RFO requests to L2 cache]
  l2_rqsts.code_rd_hit                              
       [L2 cache hits when fetching instructions, code reads]
  l2_rqsts.code_rd_miss                             
       [L2 cache misses when fetching instructions]
  l2_rqsts.demand_data_rd_hit                       
       [Demand Data Read requests that hit L2 cache]
  l2_rqsts.demand_data_rd_miss                      
       [Demand Data Read miss L2, no rejects]
  l2_rqsts.rfo_hit                                  
       [RFO requests that hit L2 cache]
  l2_rqsts.rfo_miss                                 
       [RFO requests that miss L2 cache]
  l2_rqsts.swpf_hit                                 
       [SW prefetch requests that hit L2 cache]
  l2_rqsts.swpf_miss                                
       [SW prefetch requests that miss L2 cache]
  l2_trans.l2_wb                                    
       [L2 writebacks that access L2 cache]
  longest_lat_cache.miss                            
       [Core-originated cacheable demand requests missed L3]
  mem_inst_retired.all_loads                        
       [All retired load instructions Supports address when precise (Precise
        event)]
  mem_inst_retired.all_stores                       
       [All retired store instructions Supports address when precise (Precise
        event)]
  mem_inst_retired.lock_loads                       
       [Retired load instructions with locked access Supports address when
        precise (Precise event)]
  mem_inst_retired.split_loads                      
       [Retired load instructions that split across a cacheline boundary
        Supports address when precise (Precise event)]
  mem_inst_retired.split_stores                     
       [Retired store instructions that split across a cacheline boundary
        Supports address when precise (Precise event)]
  mem_inst_retired.stlb_miss_loads                  
       [Retired load instructions that miss the STLB Supports address when
        precise (Precise event)]
  mem_inst_retired.stlb_miss_stores                 
       [Retired store instructions that miss the STLB Supports address when
        precise (Precise event)]
  mem_load_l3_hit_retired.xsnp_fwd                  
       [Retired load instructions whose data sources were HitM responses from
        shared L3 Supports address when precise (Precise event)]
  mem_load_l3_hit_retired.xsnp_hit                  
       [This event is deprecated. Refer to new event
        MEM_LOAD_L3_HIT_RETIRED.XSNP_NO_FWD Supports address when precise
        (Precise event)]
  mem_load_l3_hit_retired.xsnp_hitm                 
       [This event is deprecated. Refer to new event
        MEM_LOAD_L3_HIT_RETIRED.XSNP_FWD Supports address when precise
        (Precise event)]
  mem_load_l3_hit_retired.xsnp_miss                 
       [Retired load instructions whose data sources were L3 hit and
        cross-core snoop missed in on-pkg core cache Supports address when
        precise (Precise event)]
  mem_load_l3_hit_retired.xsnp_no_fwd               
       [Retired load instructions whose data sources were L3 and cross-core
        snoop hits in on-pkg core cache Supports address when precise (Precise
        event)]
  mem_load_l3_hit_retired.xsnp_none                 
       [Retired load instructions whose data sources were hits in L3 without
        snoops required Supports address when precise (Precise event)]
  mem_load_l3_miss_retired.local_dram               
       [Retired load instructions which data sources missed L3 but serviced
        from local dram Supports address when precise (Precise event)]
  mem_load_l3_miss_retired.remote_dram              
       [Retired load instructions which data sources missed L3 but serviced
        from remote dram Supports address when precise (Precise event)]
  mem_load_l3_miss_retired.remote_fwd               
       [Retired load instructions whose data sources was forwarded from a
        remote cache Supports address when precise (Precise event)]
  mem_load_l3_miss_retired.remote_hitm              
       [Retired load instructions whose data sources was remote HITM Supports
        address when precise (Precise event)]
  mem_load_l3_miss_retired.remote_pmm               
       [Retired demand load instructions which missed L3 but serviced from
        remote IXP memory as data sources Supports address when precise
        (Precise event)]
  mem_load_retired.fb_hit                           
       [Number of completed demand load requests that missed the L1, but hit
        the FB(fill buffer), because a preceding miss to the same cacheline
        initiated the line to be brought into L1, but data is not yet ready in
        L1 Supports address when precise (Precise event)]
  mem_load_retired.l1_hit                           
       [Retired load instructions with L1 cache hits as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.l1_miss                          
       [Retired load instructions missed L1 cache as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.l2_hit                           
       [Retired load instructions with L2 cache hits as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.l2_miss                          
       [Retired load instructions missed L2 cache as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.l3_hit                           
       [Retired load instructions with L3 cache hits as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.l3_miss                          
       [Retired load instructions missed L3 cache as data sources Supports
        address when precise (Precise event)]
  mem_load_retired.local_pmm                        
       [Retired demand load instructions which missed L3 but serviced from
        local IXP memory as data sources Supports address when precise
        (Precise event)]
  offcore_requests.all_data_rd                      
       [Demand and prefetch data reads]
  offcore_requests.all_requests                     
       [Counts memory transactions sent to the uncore]
  offcore_requests.demand_data_rd                   
       [Demand Data Read requests sent to uncore]
  offcore_requests_outstanding.all_data_rd          
       [For every cycle, increments by the number of outstanding data read
        requests the core is waiting on]
  offcore_requests_outstanding.cycles_with_data_rd  
       [For every cycle where the core is waiting on at least 1 outstanding
        demand data read request, increments by 1]
  offcore_requests_outstanding.cycles_with_demand_rfo
       [For every cycle where the core is waiting on at least 1 outstanding
        Demand RFO request, increments by 1]
  sq_misc.sq_full                                   
       [Cycles the queue waiting for offcore responses is full]

floating point:
  assists.fp                                        
       [Counts all microcode FP assists]
  fp_arith_inst_retired.128b_packed_double          
       [Counts number of SSE/AVX computational 128-bit packed double precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 2 computation operations,
        one for each element. Applies to SSE* and AVX* packed double precision
        floating-point instructions: ADD SUB HADD HSUB SUBADD MUL DIV MIN MAX
        SQRT DPP FM(N)ADD/SUB. DPP and FM(N)ADD/SUB instructions count twice
        as they perform 2 calculations per element]
  fp_arith_inst_retired.128b_packed_single          
       [Number of SSE/AVX computational 128-bit packed single precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 4 computation operations,
        one for each element. Applies to SSE* and AVX* packed single precision
        floating-point instructions: ADD SUB MUL DIV MIN MAX RCP14 RSQRT14
        SQRT DPP FM(N)ADD/SUB. DPP and FM(N)ADD/SUB instructions count twice
        as they perform 2 calculations per element]
  fp_arith_inst_retired.256b_packed_double          
       [Counts number of SSE/AVX computational 256-bit packed double precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 4 computation operations,
        one for each element. Applies to SSE* and AVX* packed double precision
        floating-point instructions: ADD SUB HADD HSUB SUBADD MUL DIV MIN MAX
        SQRT FM(N)ADD/SUB. FM(N)ADD/SUB instructions count twice as they
        perform 2 calculations per element]
  fp_arith_inst_retired.256b_packed_single          
       [Counts number of SSE/AVX computational 256-bit packed single precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 8 computation operations,
        one for each element. Applies to SSE* and AVX* packed single precision
        floating-point instructions: ADD SUB HADD HSUB SUBADD MUL DIV MIN MAX
        SQRT RSQRT RCP DPP FM(N)ADD/SUB. DPP and FM(N)ADD/SUB instructions
        count twice as they perform 2 calculations per element]
  fp_arith_inst_retired.512b_packed_double          
       [Counts number of SSE/AVX computational 512-bit packed double precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 8 computation operations,
        one for each element. Applies to SSE* and AVX* packed double precision
        floating-point instructions: ADD SUB MUL DIV MIN MAX SQRT RSQRT14
        RCP14 FM(N)ADD/SUB. FM(N)ADD/SUB instructions count twice as they
        perform 2 calculations per element]
  fp_arith_inst_retired.512b_packed_single          
       [Counts number of SSE/AVX computational 512-bit packed double precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 16 computation operations,
        one for each element. Applies to SSE* and AVX* packed double precision
        floating-point instructions: ADD SUB MUL DIV MIN MAX SQRT RSQRT14
        RCP14 FM(N)ADD/SUB. FM(N)ADD/SUB instructions count twice as they
        perform 2 calculations per element]
  fp_arith_inst_retired.scalar_double               
       [Counts number of SSE/AVX computational scalar double precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 1 computational operation.
        Applies to SSE* and AVX* scalar double precision floating-point
        instructions: ADD SUB MUL DIV MIN MAX SQRT FM(N)ADD/SUB. FM(N)ADD/SUB
        instructions count twice as they perform 2 calculations per element]
  fp_arith_inst_retired.scalar_single               
       [Counts number of SSE/AVX computational scalar single precision
        floating-point instructions retired; some instructions will count
        twice as noted below. Each count represents 1 computational operation.
        Applies to SSE* and AVX* scalar single precision floating-point
        instructions: ADD SUB MUL DIV MIN MAX SQRT RSQRT RCP FM(N)ADD/SUB.
        FM(N)ADD/SUB instructions count twice as they perform 2 calculations
        per element]

frontend:
  baclears.any                                      
       [Counts the total number when the front end is resteered, mainly when
        the BPU cannot provide a correct prediction and this is corrected by
        other branch handling mechanisms at the front end]
  dsb2mite_switches.count                           
       [Decode Stream Buffer (DSB)-to-MITE transitions count]
  dsb2mite_switches.penalty_cycles                  
       [DSB-to-MITE switch true penalty cycles]
  frontend_retired.dsb_miss                         
       [Retired Instructions who experienced DSB miss (Precise event)]
  frontend_retired.itlb_miss                        
       [Retired Instructions who experienced iTLB true miss (Precise event)]
  frontend_retired.l1i_miss                         
       [Retired Instructions who experienced Instruction L1 Cache true miss
        (Precise event)]
  frontend_retired.l2_miss                          
       [Retired Instructions who experienced Instruction L2 Cache true miss
        (Precise event)]
  frontend_retired.latency_ge_1                     
       [Retired instructions after front-end starvation of at least 1 cycle
        (Precise event)]
  frontend_retired.latency_ge_128                   
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 128 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_16                    
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 16 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_2                     
       [Retired instructions after front-end starvation of at least 2 cycles
        (Precise event)]
  frontend_retired.latency_ge_256                   
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 256 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_2_bubbles_ge_1        
       [Retired instructions that are fetched after an interval where the
        front-end had at least 1 bubble-slot for a period of 2 cycles which
        was not interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_32                    
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 32 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_4                     
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 4 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_512                   
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 512 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_64                    
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 64 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.latency_ge_8                     
       [Retired instructions that are fetched after an interval where the
        front-end delivered no uops for a period of 8 cycles which was not
        interrupted by a back-end stall (Precise event)]
  frontend_retired.stlb_miss                        
       [Retired Instructions who experienced STLB (2nd level TLB) true miss
        (Precise event)]
  icache_16b.ifdata_stall                           
       [Cycles where a code fetch is stalled due to L1 instruction cache miss]
  icache_64b.iftag_hit                              
       [Instruction fetch tag lookups that hit in the instruction cache (L1I).
        Counts at 64-byte cache-line granularity]
  icache_64b.iftag_miss                             
       [Instruction fetch tag lookups that miss in the instruction cache
        (L1I). Counts at 64-byte cache-line granularity]
  icache_64b.iftag_stall                            
       [Cycles where a code fetch is stalled due to L1 instruction cache tag
        miss]
  idq.dsb_cycles_any                                
       [Cycles Decode Stream Buffer (DSB) is delivering any Uop]
  idq.dsb_cycles_ok                                 
       [Cycles DSB is delivering optimal number of Uops]
  idq.dsb_uops                                      
       [Uops delivered to Instruction Decode Queue (IDQ) from the Decode
        Stream Buffer (DSB) path]
  idq.mite_cycles_any                               
       [Cycles MITE is delivering any Uop]
  idq.mite_cycles_ok                                
       [Cycles MITE is delivering optimal number of Uops]
  idq.mite_uops                                     
       [Uops delivered to Instruction Decode Queue (IDQ) from MITE path]
  idq.ms_switches                                   
       [Number of switches from DSB or MITE to the MS]
  idq.ms_uops                                       
       [Uops delivered to IDQ while MS is busy]
  idq_uops_not_delivered.core                       
       [Uops not delivered by IDQ when backend of the machine is not stalled]
  idq_uops_not_delivered.cycles_0_uops_deliv.core   
       [Cycles when no uops are not delivered by the IDQ when backend of the
        machine is not stalled]
  idq_uops_not_delivered.cycles_fe_was_ok           
       [Cycles when optimal number of uops was delivered to the back-end when
        the back-end is not stalled]

memory:
  cycle_activity.stalls_l3_miss                     
       [Execution stalls while L3 cache miss demand load is outstanding]
  machine_clears.memory_ordering                    
       [Number of machine clears due to memory ordering conflicts]
  mem_trans_retired.load_latency_gt_128             
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 128 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_16              
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 16 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_256             
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 256 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_32              
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 32 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_4               
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 4 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_512             
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 512 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_64              
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 64 cycles Supports address when precise
        (Must be precise)]
  mem_trans_retired.load_latency_gt_8               
       [Counts randomly selected loads when the latency from first dispatch to
        completion is greater than 8 cycles Supports address when precise
        (Must be precise)]
  rtm_retired.aborted                               
       [Number of times an RTM execution aborted]
  rtm_retired.aborted_events                        
       [Number of times an RTM execution aborted due to none of the previous 4
        categories (e.g. interrupt)]
  rtm_retired.aborted_mem                           
       [Number of times an RTM execution aborted due to various memory events
        (e.g. read/write capacity and conflicts)]
  rtm_retired.aborted_memtype                       
       [Number of times an RTM execution aborted due to incompatible memory
        type]
  rtm_retired.aborted_unfriendly                    
       [Number of times an RTM execution aborted due to HLE-unfriendly
        instructions]
  rtm_retired.commit                                
       [Number of times an RTM execution successfully committed]
  rtm_retired.start                                 
       [Number of times an RTM execution started]
  tx_exec.misc2                                     
       [Counts the number of times a class of instructions that may cause a
        transactional abort was executed inside a transactional region]
  tx_exec.misc3                                     
       [Number of times an instruction execution caused the transactional nest
        count supported to be exceeded]
  tx_mem.abort_capacity_read                        
       [Speculatively counts the number of TSX aborts due to a data capacity
        limitation for transactional reads]
  tx_mem.abort_capacity_write                       
       [Speculatively counts the number of TSX aborts due to a data capacity
        limitation for transactional writes]
  tx_mem.abort_conflict                             
       [Number of times a transactional abort was signaled due to a data
        conflict on a transactionally accessed address]

other:
  assists.any                                       
       [Number of occurrences where a microcode assist is invoked by hardware]
  core_power.lvl0_turbo_license                     
       [Core cycles where the core was running in a manner where Turbo may be
        clipped to the Non-AVX turbo schedule]
  core_power.lvl1_turbo_license                     
       [Core cycles where the core was running in a manner where Turbo may be
        clipped to the AVX2 turbo schedule]
  core_power.lvl2_turbo_license                     
       [Core cycles where the core was running in a manner where Turbo may be
        clipped to the AVX512 turbo schedule]
  ocr.demand_data_rd.l3_hit.snoop_hit_with_fwd      
       [Counts demand data reads that hit a cacheline in the L3 where a snoop
        hit in another cores caches which forwarded the unmodified data to the
        requesting core]
  ocr.demand_data_rd.l3_hit.snoop_hitm              
       [Counts demand data reads that hit a cacheline in the L3 where a snoop
        hit in another cores caches, data forwarding is required as the data
        is modified]
  ocr.demand_rfo.l3_hit.snoop_hitm                  
       [Counts writes that generate a demand reads for ownership (RFO) request
        and software prefetches for exclusive ownership (PREFETCHW) that hit a
        cacheline in the L3 where a snoop hit in another cores caches, data
        forwarding is required as the data is modified]
  ocr.streaming_wr.any_response                     
       [Counts streaming stores that have any type of response]
  sw_prefetch_access.nta                            
       [Number of PREFETCHNTA instructions executed]
  sw_prefetch_access.prefetchw                      
       [Number of PREFETCHW instructions executed]
  sw_prefetch_access.t0                             
       [Number of PREFETCHT0 instructions executed]
  sw_prefetch_access.t1_t2                          
       [Number of PREFETCHT1 or PREFETCHT2 instructions executed]
  topdown.backend_bound_slots                       
       [TMA slots where no uops were being issued due to lack of back-end
        resources]
  topdown.slots                                     
       [TMA slots available for an unhalted logical processor. Fixed counter -
        architectural event]
  topdown.slots_p                                   
       [TMA slots available for an unhalted logical processor. General counter
        - architectural event]

pipeline:
  arith.divider_active                              
       [Cycles when divide unit is busy executing divide or square root
        operations]
  br_inst_retired.all_branches                      
       [All branch instructions retired (Precise event)]
  br_inst_retired.cond                              
       [Conditional branch instructions retired (Precise event)]
  br_inst_retired.cond_ntaken                       
       [Not taken branch instructions retired (Precise event)]
  br_inst_retired.cond_taken                        
       [Taken conditional branch instructions retired (Precise event)]
  br_inst_retired.far_branch                        
       [Far branch instructions retired (Precise event)]
  br_inst_retired.indirect                          
       [All indirect branch instructions retired (excluding RETs. TSX aborts
        are considered indirect branch) (Precise event)]
  br_inst_retired.near_call                         
       [Direct and indirect near call instructions retired (Precise event)]
  br_inst_retired.near_return                       
       [Return instructions retired (Precise event)]
  br_inst_retired.near_taken                        
       [Taken branch instructions retired (Precise event)]
  br_misp_retired.all_branches                      
       [All mispredicted branch instructions retired (Precise event)]
  br_misp_retired.cond                              
       [Mispredicted conditional branch instructions retired (Precise event)]
  br_misp_retired.cond_ntaken                       
       [Mispredicted non-taken conditional branch instructions retired
        (Precise event)]
  br_misp_retired.cond_taken                        
       [number of branch instructions retired that were mispredicted and
        taken. Non PEBS (Precise event)]
  br_misp_retired.indirect                          
       [All miss-predicted indirect branch instructions retired (excluding
        RETs. TSX aborts is considered indirect branch) (Precise event)]
  br_misp_retired.near_taken                        
       [Number of near branch instructions retired that were mispredicted and
        taken (Precise event)]
  cpu_clk_unhalted.distributed                      
       [Cycle counts are evenly distributed between active threads in the Core]
  cpu_clk_unhalted.one_thread_active                
       [Core crystal clock cycles when this thread is unhalted and the other
        thread is halted]
  cpu_clk_unhalted.ref_distributed                  
       [Core crystal clock cycles. Cycle counts are evenly distributed between
        active threads in the Core]
  cpu_clk_unhalted.ref_tsc                          
       [Reference cycles when the core is not in halt state]
  cpu_clk_unhalted.ref_xclk                         
       [Core crystal clock cycles when the thread is unhalted]
  cpu_clk_unhalted.thread                           
       [Core cycles when the thread is not in halt state]
  cpu_clk_unhalted.thread_p                         
       [Thread cycles when thread is not in halt state]
  cycle_activity.cycles_l1d_miss                    
       [Cycles while L1 cache miss demand load is outstanding]
  cycle_activity.cycles_l2_miss                     
       [Cycles while L2 cache miss demand load is outstanding]
  cycle_activity.cycles_mem_any                     
       [Cycles while memory subsystem has an outstanding load]
  cycle_activity.stalls_l1d_miss                    
       [Execution stalls while L1 cache miss demand load is outstanding]
  cycle_activity.stalls_l2_miss                     
       [Execution stalls while L2 cache miss demand load is outstanding]
  cycle_activity.stalls_mem_any                     
       [Execution stalls while memory subsystem has an outstanding load]
  cycle_activity.stalls_total                       
       [Total execution stalls]
  exe_activity.1_ports_util                         
       [Cycles total of 1 uop is executed on all ports and Reservation Station
        was not empty]
  exe_activity.2_ports_util                         
       [Cycles total of 2 uops are executed on all ports and Reservation
        Station was not empty]
  exe_activity.3_ports_util                         
       [Cycles total of 3 uops are executed on all ports and Reservation
        Station was not empty]
  exe_activity.4_ports_util                         
       [Cycles total of 4 uops are executed on all ports and Reservation
        Station was not empty]
  exe_activity.bound_on_stores                      
       [Cycles where the Store Buffer was full and no loads caused an
        execution stall]
  ild_stall.lcp                                     
       [Stalls caused by changing prefix length of the instruction]
  inst_retired.any                                  
       [Number of instructions retired. Fixed Counter - architectural event
        (Precise event)]
  inst_retired.any_p                                
       [Number of instructions retired. General Counter - architectural event
        (Precise event)]
  inst_retired.prec_dist                            
       [Precise instruction retired event with a reduced effect of PEBS shadow
        in IP distribution (Precise event)]
  int_misc.all_recovery_cycles                      
       [Cycles the Backend cluster is recovering after a miss-speculation or a
        Store Buffer or Load Buffer drain stall]
  int_misc.clear_resteer_cycles                     
       [Counts cycles after recovery from a branch misprediction or machine
        clear till the first uop is issued from the resteered path]
  int_misc.recovery_cycles                          
       [Core cycles the allocator was stalled due to recovery from earlier
        clear event for this thread]
  int_misc.uop_dropping                             
       [TMA slots where uops got dropped]
  ld_blocks.no_sr                                   
       [The number of times that split load operations are temporarily blocked
        because all resources for handling the split accesses are in use]
  ld_blocks.store_forward                           
       [Loads blocked due to overlapping with a preceding store that cannot be
        forwarded]
  ld_blocks_partial.address_alias                   
       [False dependencies due to partial compare on address]
  load_hit_prefetch.swpf                            
       [Counts the number of demand load dispatches that hit L1D fill buffer
        (FB) allocated for software prefetch]
  lsd.cycles_active                                 
       [Cycles Uops delivered by the LSD, but didn't come from the decoder]
  lsd.cycles_ok                                     
       [Cycles optimal number of Uops delivered by the LSD, but did not come
        from the decoder]
  lsd.uops                                          
       [Number of Uops delivered by the LSD]
  machine_clears.count                              
       [Number of machine clears (nukes) of any type]
  machine_clears.smc                                
       [Self-modifying code (SMC) detected]
  misc_retired.pause_inst                           
       [Number of retired PAUSE instructions. This event is not supported on
        first SKL and KBL products]
  resource_stalls.sb                                
       [Cycles stalled due to no store buffers available. (not including
        draining form sync)]
  resource_stalls.scoreboard                        
       [Counts cycles where the pipeline is stalled due to serializing
        operations]
  rs_events.empty_cycles                            
       [Cycles when Reservation Station (RS) is empty for the thread]
  rs_events.empty_end                               
       [Counts end of periods where the Reservation Station (RS) was empty]
  uops_dispatched.port_0                            
       [Number of uops executed on port 0]
  uops_dispatched.port_1                            
       [Number of uops executed on port 1]
  uops_dispatched.port_2_3                          
       [Number of uops executed on port 2 and 3]
  uops_dispatched.port_4_9                          
       [Number of uops executed on port 4 and 9]
  uops_dispatched.port_5                            
       [Number of uops executed on port 5]
  uops_dispatched.port_6                            
       [Number of uops executed on port 6]
  uops_dispatched.port_7_8                          
       [Number of uops executed on port 7 and 8]
  uops_executed.core_cycles_ge_1                    
       [Cycles at least 1 micro-op is executed from any thread on physical
        core]
  uops_executed.core_cycles_ge_2                    
       [Cycles at least 2 micro-op is executed from any thread on physical
        core]
  uops_executed.core_cycles_ge_3                    
       [Cycles at least 3 micro-op is executed from any thread on physical
        core]
  uops_executed.core_cycles_ge_4                    
       [Cycles at least 4 micro-op is executed from any thread on physical
        core]
  uops_executed.cycles_ge_1                         
       [Cycles where at least 1 uop was executed per-thread]
  uops_executed.cycles_ge_2                         
       [Cycles where at least 2 uops were executed per-thread]
  uops_executed.cycles_ge_3                         
       [Cycles where at least 3 uops were executed per-thread]
  uops_executed.cycles_ge_4                         
       [Cycles where at least 4 uops were executed per-thread]
  uops_executed.stall_cycles                        
       [Counts number of cycles no uops were dispatched to be executed on this
        thread]
  uops_executed.thread                              
       [Counts the number of uops to be executed per-thread each cycle]
  uops_executed.x87                                 
       [Counts the number of x87 uops dispatched]
  uops_issued.any                                   
       [Uops that RAT issues to RS]
  uops_issued.stall_cycles                          
       [Cycles when RAT does not issue Uops to RS for the thread]
  uops_issued.vector_width_mismatch                 
       [Uops inserted at issue-stage in order to preserve upper bits of vector
        registers]
  uops_retired.slots                                
       [Retirement slots used]
  uops_retired.total_cycles                         
       [Cycles with less than 10 actually retired uops]

uncore memory:
  unc_m_act_count.all                               
       [DRAM Activate Count : All Activates. Unit: uncore_imc]
  unc_m_cas_count.all                               
       [All DRAM CAS commands issued. Unit: uncore_imc]
  unc_m_cas_count.rd                                
       [All DRAM read CAS commands issued (including underfills). Unit:
        uncore_imc]
  unc_m_cas_count.wr                                
       [All DRAM write CAS commands issued. Unit: uncore_imc]
  unc_m_clockticks                                  
       [DRAM Clockticks. Unit: uncore_imc]
  unc_m_dram_refresh.high                           
       [Number of DRAM Refreshes Issued. Unit: uncore_imc]
  unc_m_dram_refresh.opportunistic                  
       [Number of DRAM Refreshes Issued. Unit: uncore_imc]
  unc_m_dram_refresh.panic                          
       [Number of DRAM Refreshes Issued. Unit: uncore_imc]
  unc_m_hclockticks                                 
       [Half clockticks for IMC. Unit: uncore_imc]
  unc_m_pmm_cmd1.all                                
       [PMM Commands : All. Unit: uncore_imc]
  unc_m_pmm_cmd1.rd                                 
       [PMM Commands : Reads - RPQ. Unit: uncore_imc]
  unc_m_pmm_cmd1.ufill_rd                           
       [PMM Commands : Underfill reads. Unit: uncore_imc]
  unc_m_pmm_cmd1.wr                                 
       [PMM Commands : Writes. Unit: uncore_imc]
  unc_m_pmm_rpq_inserts                             
       [PMM Read Queue Inserts. Unit: uncore_imc]
  unc_m_pmm_rpq_occupancy.all                       
       [PMM Read Pending Queue Occupancy. Unit: uncore_imc]
  unc_m_pmm_wpq_inserts                             
       [PMM Write Queue Inserts. Unit: uncore_imc]
  unc_m_pmm_wpq_occupancy.all                       
       [PMM Write Pending Queue Occupancy. Unit: uncore_imc]
  unc_m_pre_count.all                               
       [DRAM Precharge commands. Unit: uncore_imc]
  unc_m_pre_count.pgt                               
       [DRAM Precharge commands. : Precharge due to page table. Unit:
        uncore_imc]
  unc_m_pre_count.rd                                
       [DRAM Precharge commands. : Precharge due to read. Unit: uncore_imc]
  unc_m_pre_count.wr                                
       [DRAM Precharge commands. : Precharge due to write. Unit: uncore_imc]
  unc_m_rpq_inserts.pch0                            
       [Read Pending Queue Allocations. Unit: uncore_imc]
  unc_m_rpq_inserts.pch1                            
       [Read Pending Queue Allocations. Unit: uncore_imc]
  unc_m_rpq_occupancy_pch0                          
       [Read Pending Queue Occupancy. Unit: uncore_imc]
  unc_m_rpq_occupancy_pch1                          
       [Read Pending Queue Occupancy. Unit: uncore_imc]
  unc_m_tagchk.hit                                  
       [2LM Tag Check : Hit in Near Memory Cache. Unit: uncore_imc]
  unc_m_tagchk.miss_clean                           
       [2LM Tag Check : Miss, no data in this line. Unit: uncore_imc]
  unc_m_tagchk.miss_dirty                           
       [2LM Tag Check : Miss, existing data may be evicted to Far Memory.
        Unit: uncore_imc]
  unc_m_tagchk.nm_rd_hit                            
       [2LM Tag Check : Read Hit in Near Memory Cache. Unit: uncore_imc]
  unc_m_tagchk.nm_wr_hit                            
       [2LM Tag Check : Write Hit in Near Memory Cache. Unit: uncore_imc]
  unc_m_wpq_inserts.pch0                            
       [Write Pending Queue Allocations. Unit: uncore_imc]
  unc_m_wpq_inserts.pch1                            
       [Write Pending Queue Allocations. Unit: uncore_imc]
  unc_m_wpq_occupancy_pch0                          
       [Write Pending Queue Occupancy. Unit: uncore_imc]
  unc_m_wpq_occupancy_pch1                          
       [Write Pending Queue Occupancy. Unit: uncore_imc]

uncore other:
  unc_cha_clockticks                                
       [Clockticks of the uncore caching &amp;amp; home agent (CHA). Unit:
        uncore_cha]
  unc_cha_cms_clockticks                            
       [CMS Clockticks. Unit: uncore_cha]
  unc_cha_imc_reads_count.normal                    
       [Normal priority reads issued to the memory controller from the CHA.
        Unit: uncore_cha]
  unc_cha_imc_writes_count.full                     
       [CHA to iMC Full Line Writes Issued : Full Line Non-ISOCH. Unit:
        uncore_cha]
  unc_cha_llc_lookup.data_read                      
       [Cache and Snoop Filter Lookups; Data Read Request. Unit: uncore_cha]
  unc_cha_llc_victims.all                           
       [Lines Victimized : All Lines Victimized. Unit: uncore_cha]
  unc_cha_requests.invitoe_local                    
       [Local INVITOE requests (exclusive ownership of a cache line without
        receiving data) that miss the SF/LLC and are sent to the CHA's home
        agent. Unit: uncore_cha]
  unc_cha_requests.invitoe_remote                   
       [Remote INVITOE requests (exclusive ownership of a cache line without
        receiving data) sent to the CHA's home agent. Unit: uncore_cha]
  unc_cha_requests.reads                            
       [Local read requests that miss the SF/LLC and remote read requests sent
        to the CHA's home agent. Unit: uncore_cha]
  unc_cha_requests.reads_local                      
       [Local read requests that miss the SF/LLC and are sent to the CHA's
        home agent. Unit: uncore_cha]
  unc_cha_requests.reads_remote                     
       [Remote read requests sent to the CHA's home agent. Unit: uncore_cha]
  unc_cha_requests.writes                           
       [Local write requests that miss the SF/LLC and remote write requests
        sent to the CHA's home agent. Unit: uncore_cha]
  unc_cha_requests.writes_local                     
       [Local write requests that miss the SF/LLC and are sent to the CHA's
        home agent. Unit: uncore_cha]
  unc_cha_requests.writes_remote                    
       [Remote write requests sent to the CHA's home agent. Unit: uncore_cha]
  unc_cha_sf_eviction.e_state                       
       [Snoop filter capacity evictions for E-state entries. Unit: uncore_cha]
  unc_cha_sf_eviction.m_state                       
       [Snoop filter capacity evictions for M-state entries. Unit: uncore_cha]
  unc_cha_sf_eviction.s_state                       
       [Snoop filter capacity evictions for S-state entries. Unit: uncore_cha]
  unc_cha_tor_inserts.ia                            
       [TOR Inserts : All requests from iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_clflush                    
       [TOR Inserts : CLFlushes issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_crd                        
       [TOR Inserts : CRDs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_drd_pref                   
       [TOR Inserts : DRd_Prefs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_hit                        
       [TOR Inserts : All requests from iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_crd                    
       [TOR Inserts : CRds issued by iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_crd_pref               
       [TOR Inserts : CRd_Prefs issued by iA Cores that hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_drd                    
       [TOR Inserts : DRds issued by iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_drd_pref               
       [TOR Inserts : DRd_Prefs issued by iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_llcprefrfo             
       [TOR Inserts : LLCPrefRFO issued by iA Cores that hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_rfo                    
       [TOR Inserts : RFOs issued by iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_hit_rfo_pref               
       [TOR Inserts : RFO_Prefs issued by iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_llcprefdata                
       [TOR Inserts : LLCPrefData issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_llcprefrfo                 
       [TOR Inserts : LLCPrefRFO issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss                       
       [TOR Inserts : All requests from iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_crd                   
       [TOR Inserts : CRds issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_crd_pref              
       [TOR Inserts : CRd_Prefs issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd                   
       [TOR Inserts : DRds issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_ddr               
       [TOR Inserts : DRds issued by iA Cores targeting DDR Mem that Missed
        the LLC. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_local             
       [TOR Inserts : DRds issued by iA Cores that Missed the LLC - HOMed
        locally. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_local_ddr         
       [TOR Inserts : DRds issued by iA Cores targeting DDR Mem that Missed
        the LLC - HOMed locally. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_local_pmm         
       [TOR Inserts : DRds issued by iA Cores targeting PMM Mem that Missed
        the LLC - HOMed locally. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_pmm               
       [TOR Inserts : DRds issued by iA Cores targeting PMM Mem that Missed
        the LLC. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_pref              
       [TOR Inserts : DRd_Prefs issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_pref_local        
       [TOR Inserts; DRd Pref misses from local IA. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_pref_remote       
       [TOR Inserts; DRd Pref misses from local IA. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_remote            
       [TOR Inserts : DRds issued by iA Cores that Missed the LLC - HOMed
        remotely. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_remote_ddr        
       [TOR Inserts : DRds issued by iA Cores targeting DDR Mem that Missed
        the LLC - HOMed remotely. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_drd_remote_pmm        
       [TOR Inserts : DRds issued by iA Cores targeting PMM Mem that Missed
        the LLC - HOMed remotely. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_full_streaming_wr     
       [TOR Inserts; WCiLF misses from local IA. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_llcprefdata           
       [TOR Inserts : LLCPrefData issued by iA Cores that missed the LLC.
        Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_llcprefrfo            
       [TOR Inserts : LLCPrefRFO issued by iA Cores that missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_partial_streaming_wr  
       [TOR Inserts; WCiL misses from local IA. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo                   
       [TOR Inserts : RFOs issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo_local             
       [TOR Inserts : RFOs issued by iA Cores that Missed the LLC - HOMed
        locally. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo_pref              
       [TOR Inserts : RFO_Prefs issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo_pref_local        
       [TOR Inserts : RFO_Prefs issued by iA Cores that Missed the LLC - HOMed
        locally. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo_pref_remote       
       [TOR Inserts : RFO_Prefs issued by iA Cores that Missed the LLC - HOMed
        remotely. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_miss_rfo_remote            
       [TOR Inserts : RFOs issued by iA Cores that Missed the LLC - HOMed
        remotely. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_rfo                        
       [TOR Inserts : RFOs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_rfo_pref                   
       [TOR Inserts : RFO_Prefs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.ia_specitom                   
       [TOR Inserts : SpecItoMs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_inserts.io                            
       [TOR Inserts : All requests from IO Devices. Unit: uncore_cha]
  unc_cha_tor_inserts.io_hit                        
       [TOR Inserts : All requests from IO Devices that hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.io_hit_itom                   
       [TOR Inserts : ItoMs issued by IO Devices that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.io_hit_itomcachenear          
       [TOR Inserts : ItoMCacheNears, indicating a partial write request, from
        IO Devices that hit the LLC. Unit: uncore_cha]
  unc_cha_tor_inserts.io_hit_pcirdcur               
       [TOR Inserts : PCIRdCurs issued by IO Devices that hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.io_itom                       
       [TOR Inserts : ItoMs issued by IO Devices. Unit: uncore_cha]
  unc_cha_tor_inserts.io_itomcachenear              
       [TOR Inserts : ItoMCacheNears, indicating a partial write request, from
        IO Devices. Unit: uncore_cha]
  unc_cha_tor_inserts.io_miss                       
       [TOR Inserts : All requests from IO Devices that missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.io_miss_itom                  
       [TOR Inserts : ItoMs issued by IO Devices that missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_inserts.io_miss_itomcachenear         
       [TOR Inserts : ItoMCacheNears, indicating a partial write request, from
        IO Devices that missed the LLC. Unit: uncore_cha]
  unc_cha_tor_inserts.io_miss_pcirdcur              
       [TOR Inserts : PCIRdCurs issued by IO Devices that missed the LLC.
        Unit: uncore_cha]
  unc_cha_tor_inserts.io_pcirdcur                   
       [TOR Inserts : PCIRdCurs issued by IO Devices. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia                          
       [TOR Occupancy : All requests from iA Cores. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_crd                      
       [TOR Occupancy : CRDs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_drd                      
       [TOR Occupancy : DRds issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_hit                      
       [TOR Occupancy : All requests from iA Cores that Hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.ia_miss                     
       [TOR Occupancy : All requests from iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.ia_miss_crd                 
       [TOR Occupancy : CRds issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.ia_miss_drd                 
       [TOR Occupancy : DRds issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.ia_miss_drd_ddr             
       [TOR Occupancy : DRds issued by iA Cores targeting DDR Mem that Missed
        the LLC. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_miss_drd_local           
       [TOR Occupancy : DRds issued by iA Cores that Missed the LLC - HOMed
        locally. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_miss_drd_pmm             
       [TOR Occupancy : DRds issued by iA Cores targeting PMM Mem that Missed
        the LLC. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_miss_drd_remote          
       [TOR Occupancy : DRds issued by iA Cores that Missed the LLC - HOMed
        remotely. Unit: uncore_cha]
  unc_cha_tor_occupancy.ia_miss_rfo                 
       [TOR Occupancy : RFOs issued by iA Cores that Missed the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.ia_rfo                      
       [TOR Occupancy : RFOs issued by iA Cores. Unit: uncore_cha]
  unc_cha_tor_occupancy.io                          
       [TOR Occupancy : All requests from IO Devices. Unit: uncore_cha]
  unc_cha_tor_occupancy.io_hit                      
       [TOR Occupancy : All requests from IO Devices that hit the LLC. Unit:
        uncore_cha]
  unc_cha_tor_occupancy.io_miss                     
       [TOR Occupancy : All requests from IO Devices that missed the LLC.
        Unit: uncore_cha]
  unc_cha_tor_occupancy.io_miss_pcirdcur            
       [TOR Occupancy : PCIRdCurs issued by IO Devices that missed the LLC.
        Unit: uncore_cha]
  unc_cha_tor_occupancy.io_pcirdcur                 
       [TOR Occupancy : PCIRdCurs issued by IO Devices. Unit: uncore_cha]
  unc_i_cache_total_occupancy.mem                   
       [Total IRP occupancy of inbound read and write requests to coherent
        memory. Unit: uncore_irp]
  unc_i_clockticks                                  
       [Clockticks of the IO coherency tracker (IRP). Unit: uncore_irp]
  unc_i_coherent_ops.pcitom                         
       [PCIITOM request issued by the IRP unit to the mesh with the intention
        of writing a full cacheline. Unit: uncore_irp]
  unc_i_coherent_ops.wbmtoi                         
       [Coherent Ops : WbMtoI. Unit: uncore_irp]
  unc_i_faf_full                                    
       [FAF RF full. Unit: uncore_irp]
  unc_i_faf_inserts                                 
       [Inbound read requests received by the IRP and inserted into the FAF
        queue. Unit: uncore_irp]
  unc_i_faf_occupancy                               
       [Occupancy of the IRP FAF queue. Unit: uncore_irp]
  unc_i_faf_transactions                            
       [FAF allocation -- sent to ADQ. Unit: uncore_irp]
  unc_i_irp_all.inbound_inserts                     
       [: All Inserts Inbound (p2p + faf + cset). Unit: uncore_irp]
  unc_i_misc1.lost_fwd                              
       [Misc Events - Set 1 : Lost Forward. Unit: uncore_irp]
  unc_i_snoop_resp.all_hit_m                        
       [Responses to snoops of any type that hit M line in the IIO cache.
        Unit: uncore_irp]
  unc_i_transactions.wr_pref                        
       [Inbound write (fast path) requests received by the IRP. Unit:
        uncore_irp]
  unc_iio_clockticks                                
       [Clockticks of the integrated IO (IIO) traffic controller. Unit:
        uncore_iio]
  unc_iio_clockticks_freerun                        
       [Free running counter that increments for IIO clocktick. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.all_parts           
       [PCIe Completion Buffer Inserts of completions with data: Part 0-7.
        Unit: uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part0               
       [PCIe Completion Buffer Inserts of completions with data: Part 0. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part1               
       [PCIe Completion Buffer Inserts of completions with data: Part 1. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part2               
       [PCIe Completion Buffer Inserts of completions with data: Part 2. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part3               
       [PCIe Completion Buffer Inserts of completions with data: Part 3. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part4               
       [PCIe Completion Buffer Inserts of completions with data: Part 4. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part5               
       [PCIe Completion Buffer Inserts of completions with data: Part 5. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part6               
       [PCIe Completion Buffer Inserts of completions with data: Part 6. Unit:
        uncore_iio]
  unc_iio_comp_buf_inserts.cmpd.part7               
       [PCIe Completion Buffer Inserts of completions with data: Part 7. Unit:
        uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.all_parts         
       [PCIe Completion Buffer Occupancy of completions with data : Part 0-7.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part0             
       [PCIe Completion Buffer Occupancy of completions with data : Part 0.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part1             
       [PCIe Completion Buffer Occupancy of completions with data : Part 1.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part2             
       [PCIe Completion Buffer Occupancy of completions with data : Part 2.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part3             
       [PCIe Completion Buffer Occupancy of completions with data : Part 3.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part4             
       [PCIe Completion Buffer Occupancy of completions with data : Part 4.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part5             
       [PCIe Completion Buffer Occupancy of completions with data : Part 5.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part6             
       [PCIe Completion Buffer Occupancy of completions with data : Part 6.
        Unit: uncore_iio]
  unc_iio_comp_buf_occupancy.cmpd.part7             
       [PCIe Completion Buffer Occupancy of completions with data : Part 7.
        Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part0            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part1            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part2            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part3            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part4            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part5            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part6            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_read.part7            
       [Data requested by the CPU : Core reporting completion of Card read
        from Core DRAM. Unit: uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part0           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part1           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part2           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part3           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part4           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part5           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part6           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_by_cpu.mem_write.part7           
       [Data requested by the CPU : Core writing to Card's MMIO space. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part0                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part1                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part2                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part3                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part4                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part5                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part6                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.cmpd.part7                
       [Data requested of the CPU : CmpD - device sending completion to CPU
        request. Unit: uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part0            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part1            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part2            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part3            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part4            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part5            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part6            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_read.part7            
       [Four byte data request of the CPU : Card reading from DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part0           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part1           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part2           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part3           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part4           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part5           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part6           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_data_req_of_cpu.mem_write.part7           
       [Four byte data request of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_num_req_of_cpu.commit.all                 
       [Number requests PCIe makes of the main die : All. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part0             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part1             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part2             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part3             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part4             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part5             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part6             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_read.part7             
       [Number Transactions requested by the CPU : Core reading from Card's
        MMIO space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part0            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part1            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part2            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part3            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part4            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part5            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part6            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_by_cpu.mem_write.part7            
       [Number Transactions requested by the CPU : Core writing to Card's MMIO
        space. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part0                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part1                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part2                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part3                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part4                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part5                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part6                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.cmpd.part7                 
       [Number Transactions requested of the CPU : CmpD - device sending
        completion to CPU request. Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part0             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part1             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part2             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part3             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part4             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part5             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part6             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_read.part7             
       [Number Transactions requested of the CPU : Card reading from DRAM.
        Unit: uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part0            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part1            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part2            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part3            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part4            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part5            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part6            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_iio_txn_req_of_cpu.mem_write.part7            
       [Number Transactions requested of the CPU : Card writing to DRAM. Unit:
        uncore_iio]
  unc_m2m_clockticks                                
       [Clockticks of the mesh to memory (M2M). Unit: uncore_m2m]
  unc_m2m_cms_clockticks                            
       [CMS Clockticks. Unit: uncore_m2m]
  unc_m2m_directory_lookup.any                      
       [Multi-socket cacheline Directory Lookups : Found in any state. Unit:
        uncore_m2m]
  unc_m2m_directory_lookup.state_a                  
       [Multi-socket cacheline Directory Lookups : Found in A state. Unit:
        uncore_m2m]
  unc_m2m_directory_lookup.state_i                  
       [Multi-socket cacheline Directory Lookups : Found in I state. Unit:
        uncore_m2m]
  unc_m2m_directory_lookup.state_s                  
       [Multi-socket cacheline Directory Lookups : Found in S state. Unit:
        uncore_m2m]
  unc_m2m_imc_reads.to_pmm                          
       [M2M Reads Issued to iMC : PMM - All Channels. Unit: uncore_m2m]
  unc_m2m_imc_writes.to_pmm                         
       [M2M Writes Issued to iMC : PMM - All Channels. Unit: uncore_m2m]
  unc_m2m_tag_hit.nm_rd_hit_clean                   
       [Tag Hit : Clean NearMem Read Hit. Unit: uncore_m2m]
  unc_m2m_tag_hit.nm_rd_hit_dirty                   
       [Tag Hit : Dirty NearMem Read Hit. Unit: uncore_m2m]
  unc_m2p_clockticks                                
       [Clockticks of the mesh to PCI (M2P). Unit: uncore_m2pcie]
  unc_m2p_cms_clockticks                            
       [CMS Clockticks. Unit: uncore_m2pcie]
  unc_m3upi_clockticks                              
       [Clockticks of the mesh to UPI (M3UPI). Unit: uncore_m3upi]
  unc_u_clockticks                                  
       [Clockticks in the UBOX using a dedicated 48-bit Fixed Counter. Unit:
        uncore_ubox]
  unc_upi_clockticks                                
       [Number of kfclks. Unit: uncore_upi]
  unc_upi_l1_power_cycles                           
       [Cycles in L1. Unit: uncore_upi]
  unc_upi_rxl_flits.all_data                        
       [Valid Flits Received : All Data. Unit: uncore_upi]
  unc_upi_rxl_flits.all_null                        
       [Valid Flits Received : Null FLITs received from any slot. Unit:
        uncore_upi]
  unc_upi_rxl_flits.non_data                        
       [Valid Flits Received : All Non Data. Unit: uncore_upi]
  unc_upi_txl0p_power_cycles                        
       [Cycles in L0p. Unit: uncore_upi]
  unc_upi_txl_flits.all_data                        
       [Valid Flits Sent : All Data. Unit: uncore_upi]
  unc_upi_txl_flits.all_null                        
       [Valid Flits Sent : Null FLITs transmitted to any slot. Unit:
        uncore_upi]
  unc_upi_txl_flits.non_data                        
       [Valid Flits Sent : All Non Data. Unit: uncore_upi]

uncore power:
  unc_p_clockticks                                  
       [Clockticks of the power control unit (PCU). Unit: uncore_pcu]

virtual memory:
  dtlb_load_misses.stlb_hit                         
       [Loads that miss the DTLB and hit the STLB]
  dtlb_load_misses.walk_active                      
       [Cycles when at least one PMH is busy with a page walk for a demand
        load]
  dtlb_load_misses.walk_completed                   
       [Load miss in all TLB levels causes a page walk that completes. (All
        page sizes)]
  dtlb_load_misses.walk_completed_2m_4m             
       [Page walks completed due to a demand data load to a 2M/4M page]
  dtlb_load_misses.walk_completed_4k                
       [Page walks completed due to a demand data load to a 4K page]
  dtlb_load_misses.walk_pending                     
       [Number of page walks outstanding for a demand load in the PMH each
        cycle]
  dtlb_store_misses.stlb_hit                        
       [Stores that miss the DTLB and hit the STLB]
  dtlb_store_misses.walk_active                     
       [Cycles when at least one PMH is busy with a page walk for a store]
  dtlb_store_misses.walk_completed                  
       [Store misses in all TLB levels causes a page walk that completes. (All
        page sizes)]
  dtlb_store_misses.walk_completed_2m_4m            
       [Page walks completed due to a demand data store to a 2M/4M page]
  dtlb_store_misses.walk_completed_4k               
       [Page walks completed due to a demand data store to a 4K page]
  dtlb_store_misses.walk_pending                    
       [Number of page walks outstanding for a store in the PMH each cycle]
  itlb_misses.stlb_hit                              
       [Instruction fetch requests that miss the ITLB and hit the STLB]
  itlb_misses.walk_active                           
       [Cycles when at least one PMH is busy with a page walk for code
        (instruction fetch) request]
  itlb_misses.walk_completed                        
       [Code miss in all TLB levels causes a page walk that completes. (All
        page sizes)]
  itlb_misses.walk_completed_2m_4m                  
       [Code miss in all TLB levels causes a page walk that completes. (2M/4M)]
  itlb_misses.walk_completed_4k                     
       [Code miss in all TLB levels causes a page walk that completes. (4K)]
  itlb_misses.walk_pending                          
       [Number of page walks outstanding for an outstanding code request in
        the PMH each cycle]
  tlb_flush.dtlb_thread                             
       [DTLB flush attempts of the thread-specific entries]
  tlb_flush.stlb_any                                
       [STLB flush attempts]
  rNNN                                               [Raw hardware event descriptor]
  cpu/t1=v1[,t2=v2,t3 ...]/modifier                  [Raw hardware event descriptor]
  mem:<addr>[/len][:access]                          [Hardware breakpoint]

Metric Groups:

BrMispredicts:
  IpMispredict
       [Number of Instructions per non-speculative Branch Misprediction (JEClear)]
Branches:
  BpTkBranch
       [Branch instructions per taken branch]
  IpBranch
       [Instructions per Branch (lower number means higher occurrence rate)]
  IpCall
       [Instructions per (near) call (lower number means higher occurrence rate)]
  IpFarBranch
       [Instructions per Far Branch ( Far Branches apply upon transition from application to operating system, handling interrupts, exceptions) [lower number means higher occurrence rate]]
  IpTB
       [Instruction per taken branch]
CacheMisses:
  L1MPKI
       [L1 cache true misses per kilo instruction for retired demand loads]
  L2MPKI
       [L2 cache true misses per kilo instruction for retired demand loads]
  L2MPKI_All
       [L2 cache misses per kilo instruction for all request types (including speculative)]
  L3MPKI
       [L3 cache true misses per kilo instruction for retired demand loads]
DSB:
  DSB_Coverage
       [Fraction of Uops delivered by the DSB (aka Decoded ICache; or Uop Cache)]
FetchBW:
  DSB_Coverage
       [Fraction of Uops delivered by the DSB (aka Decoded ICache; or Uop Cache)]
  IpTB
       [Instruction per taken branch]
Flops:
  FLOPc
       [Floating Point Operations Per Cycle]
  GFLOPs
       [Giga Floating Point Operations Per Second]
  IpFLOP
       [Instructions per Floating Point (FP) Operation (lower number means higher occurrence rate)]
FpArith:
  IpFLOP
       [Instructions per Floating Point (FP) Operation (lower number means higher occurrence rate)]
HPC:
  CPU_Utilization
       [Average CPU Utilization]
  DRAM_BW_Use
       [Average external Memory Bandwidth Use for reads and writes [GB / sec]]
  GFLOPs
       [Giga Floating Point Operations Per Second]
InsType:
  IpBranch
       [Instructions per Branch (lower number means higher occurrence rate)]
  IpFLOP
       [Instructions per Floating Point (FP) Operation (lower number means higher occurrence rate)]
  IpLoad
       [Instructions per Load (lower number means higher occurrence rate)]
  IpStore
       [Instructions per Store (lower number means higher occurrence rate)]
IoBW:
  IO_Read_BW
       [Average IO (network or disk) Bandwidth Use for Reads [GB / sec]]
  IO_Write_BW
       [Average IO (network or disk) Bandwidth Use for Writes [GB / sec]]
L2Evicts:
  L2_Evictions_NonSilent_PKI
       [Rate of non silent evictions from the L2 cache per Kilo instruction]
  L2_Evictions_Silent_PKI
       [Rate of silent evictions from the L2 cache per Kilo instruction where the evicted lines are dropped (no writeback to L3 or memory)]
LSD:
  LSD_Coverage
       [Fraction of Uops delivered by the LSD (Loop Stream Detector; aka Loop Cache)]
MemoryBW:
  DRAM_BW_Use
       [Average external Memory Bandwidth Use for reads and writes [GB / sec]]
  L1D_Cache_Fill_BW
       [Average data fill bandwidth to the L1 data cache [GB / sec]]
  L2_Cache_Fill_BW
       [Average data fill bandwidth to the L2 cache [GB / sec]]
  L3_Cache_Access_BW
       [Average per-core data access bandwidth to the L3 cache [GB / sec]]
  L3_Cache_Fill_BW
       [Average per-core data fill bandwidth to the L3 cache [GB / sec]]
  MEM_Parallel_Reads
       [Average number of parallel data read requests to external memory. Accounts for demand loads and L1/L2 prefetches]
  MLP
       [Memory-Level-Parallelism (average number of L1 miss demand load when there is at least one such miss. Per-Logical Processor)]
  PMM_Read_BW
       [Average 3DXP Memory Bandwidth Use for reads [GB / sec]]
  PMM_Write_BW
       [Average 3DXP Memory Bandwidth Use for Writes [GB / sec]]
MemoryBound:
  Load_Miss_Real_Latency
       [Actual Average Latency for L1 data-cache miss demand loads (in core cycles)]
  MLP
       [Memory-Level-Parallelism (average number of L1 miss demand load when there is at least one such miss. Per-Logical Processor)]
MemoryLat:
  Load_Miss_Real_Latency
       [Actual Average Latency for L1 data-cache miss demand loads (in core cycles)]
  MEM_PMM_Read_Latency
       [Average latency of data read request to external 3D X-Point memory [in nanoseconds]. Accounts for demand loads and L1/L2 data-read prefetches]
  MEM_Read_Latency
       [Average latency of data read request to external memory (in nanoseconds). Accounts for demand loads and L1/L2 prefetches]
MemoryTLB:
  Page_Walks_Utilization
       [Utilization of the core's Page Walker(s) serving STLB misses triggered by instruction/Load/Store accesses]
OS:
  IpFarBranch
       [Instructions per Far Branch ( Far Branches apply upon transition from application to operating system, handling interrupts, exceptions) [lower number means higher occurrence rate]]
  Kernel_Utilization
       [Fraction of cycles spent in the Operating System (OS) Kernel mode]
Offcore:
  L2MPKI_All
       [L2 cache misses per kilo instruction for all request types (including speculative)]
  L3_Cache_Access_BW
       [Average per-core data access bandwidth to the L3 cache [GB / sec]]
PGO:
  BpTkBranch
       [Branch instructions per taken branch]
  IpTB
       [Instruction per taken branch]
Pipeline:
  CLKS
       [Per-Logical Processor actual clocks when the Logical Processor is active]
  CPI
       [Cycles Per Instruction (per Logical Processor)]
  ILP
       [Instruction-Level-Parallelism (average number of uops executed when there is at least 1 uop executed)]
  UPI
       [Uops Per Instruction]
PortsUtil:
  ILP
       [Instruction-Level-Parallelism (average number of uops executed when there is at least 1 uop executed)]
Power:
  Average_Frequency
       [Measured Average Frequency for unhalted processors [GHz]]
  C1_Core_Residency
       [C1 residency percent per core]
  C2_Pkg_Residency
       [C2 residency percent per package]
  C6_Core_Residency
       [C6 residency percent per core]
  C6_Pkg_Residency
       [C6 residency percent per package]
  Turbo_Utilization
       [Average Frequency Utilization relative nominal frequency]
Retire:
  UPI
       [Uops Per Instruction]
SMT:
  CORE_CLKS
       [Core actual clocks when any Logical Processor is active on the Physical Core]
  CoreIPC
       [Instructions Per Cycle (per physical core)]
  SMT_2T_Utilization
       [Fraction of cycles where both hardware Logical Processors were active]
Server:
  IO_Read_BW
       [Average IO (network or disk) Bandwidth Use for Reads [GB / sec]]
  IO_Write_BW
       [Average IO (network or disk) Bandwidth Use for Writes [GB / sec]]
  L2_Evictions_NonSilent_PKI
       [Rate of non silent evictions from the L2 cache per Kilo instruction]
  L2_Evictions_Silent_PKI
       [Rate of silent evictions from the L2 cache per Kilo instruction where the evicted lines are dropped (no writeback to L3 or memory)]
  MEM_PMM_Read_Latency
       [Average latency of data read request to external 3D X-Point memory [in nanoseconds]. Accounts for demand loads and L1/L2 data-read prefetches]
  PMM_Read_BW
       [Average 3DXP Memory Bandwidth Use for reads [GB / sec]]
  PMM_Write_BW
       [Average 3DXP Memory Bandwidth Use for Writes [GB / sec]]
SoC:
  DRAM_BW_Use
       [Average external Memory Bandwidth Use for reads and writes [GB / sec]]
  IO_Read_BW
       [Average IO (network or disk) Bandwidth Use for Reads [GB / sec]]
  IO_Write_BW
       [Average IO (network or disk) Bandwidth Use for Writes [GB / sec]]
  MEM_PMM_Read_Latency
       [Average latency of data read request to external 3D X-Point memory [in nanoseconds]. Accounts for demand loads and L1/L2 data-read prefetches]
  MEM_Parallel_Reads
       [Average number of parallel data read requests to external memory. Accounts for demand loads and L1/L2 prefetches]
  MEM_Read_Latency
       [Average latency of data read request to external memory (in nanoseconds). Accounts for demand loads and L1/L2 prefetches]
  PMM_Read_BW
       [Average 3DXP Memory Bandwidth Use for reads [GB / sec]]
  PMM_Write_BW
       [Average 3DXP Memory Bandwidth Use for Writes [GB / sec]]
  Socket_CLKS
       [Socket actual clocks when any core is active on that socket]
Summary:
  Average_Frequency
       [Measured Average Frequency for unhalted processors [GHz]]
  CPU_Utilization
       [Average CPU Utilization]
  IPC
       [Instructions Per Cycle (per Logical Processor)]
  Instructions
       [Total number of retired Instructions, Sample with: INST_RETIRED.PREC_DIST]
TmaL1:
  CoreIPC
       [Instructions Per Cycle (per physical core)]
  Instructions
       [Total number of retired Instructions, Sample with: INST_RETIRED.PREC_DIST]

