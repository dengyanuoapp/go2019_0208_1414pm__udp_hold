

CFGmakeRun:=Makefile.run.go01.mk


GoTOP:=go2019_0208_1414pm__udp_hold
GoTOP:= \
	udp_2201__Fn \
	udp_2301__Dn \
	udp_2401__Cn \

#	udp_0001_test_udp \
	#	udp_0997__server  udp_0998__client udp_0888__testOnly

GoTOP:=$(strip $(GoTOP))

#-----------====----====----
# the id must be 128 bytes : 16 byte
fn_id128:=0xFF22afd83a8a8ac98a8d8a0183c822FF
dn_id128:=0xDD33928381a88a7aa7d9a8d9182833DD
cn_id128:=0xCC44928381a88a7aa7d9a8d9182844CC
#Fn_id128:=8389afd83a8a8ac98a8d8a0183c839b1
#Dn_id128:=8381928381a88a7aa7d9a8d918289ab2
##---------====----====----====----====----


udp_0997__server:= udp_0997__server  udp_0901__type  
udp_0998__client:= udp_0998__client  udp_0901__type  
udp_0001_test_udp:= udp_0001_test_udp
udp_0888__testOnly:= 			\
	base_1121__Sprintf.go								\
	base_1122__Fprintf.go								\
	base_1123__Cprintf.go								\
 	\
	base_1103__checkError								\
	\
	base_1997__encode_decode__json						\
	base_1998__encode_decode__gob						\
	base_1999__encode_decode__bin						\
	\
	udp_0888__testOnly

udp___base:=											\
	$(patsubst %.go,%,$(shell cd src0 && ls base_1*.go))								\
	udp_1899__FnDnSnCn_const_passwd

udp_2201__Fn:=$(udp___base)	udp_2201__Fn 
udp_2301__Dn:=$(udp___base)	udp_2301__Dn
udp_2401__Cn:=$(udp___base)	udp_2401__Cn

wputPATH:=ftp://coe:coe@192.168.1.93/


DockerList:=$(GoTOP)
# Makefile.1021.docker01



dbc:=docker_build_clean
$(dbc):= dkcc dkcn dkci     cl 
dbc:$($(dbc))

dbb:=docker_build_build
$(dbb):= bl     dkb
dbb:$($(dbb))

showRunHelpListLast += dbc dbb  $(testList) 

GoPreLinuxALL:= LinuxX64 
GoPreDockerALL:=LinuxX64 

bt: btgo


f:=udp_2201__Fn  src2
d:=udp_2301__Dn  src3
c:=udp_2401__Cn  src4

testListEnv:=f d c

define testTP1

aaa$(1)1:=$(word 1,$($(1)))
aaa$(1)2:=$(word 2,$($(1)))
aaa$(1)3:=$$(aaa$(1)2)/$$(aaa$(1)1).go

bbb$(1)1:=$$(shell echo $$(aaa$(1)1)|tr -d '[\\.\-_]' |tr [A-Z] [a-z])
bbb$(1)2:=dengyanuoapp/$$(bbb$(1)1)lnxlinuxx64exe
bbb$(1)3:=lnx/$$(aaa$(1)1).lnx.LinuxX64.exe

endef
define testTP2
r$(1)1:=run_vim__$(1)
$$(r$(1)1):= make vg DST=$$(aaa$(1)3)

r$(1)2:=run_build_without_upload_docker__$(1)
r$(1)2:
	@echo
	rm -f $$(bbb$(1)3)
	make bl 
	@echo
	make noupload=1 dkb 

r$(1)3:=run_test1_only__$(1)
$$(r$(1)3):= nice -n 19 docker run  --memory-swappiness 0 -m 100m --rm -it \
	--network host -e id128=$($(1)n_id128) \
	--name run__$(1) \
	$$(bbb$(1)2)   \

r$(1)4:=commit_gen_image_for__$(1)
r$(1)4:
	-docker image rmi udp_test/$(1)n 2>/dev/null 
	docker commit \
		`docker ps -a |grep $$(bbb$(1)2) |head -n 1|awk '{print $$1}'` \
		udp_test/$(1)n

r$(1)9:=run_all__$(1)
r$(1)9: $$(r$(1)1) $$(r$(1)2) $$(r$(1)3)


testList2+= r$(1)1 r$(1)2 r$(1)3 r$(1)4 r$(1)9 

endef

define testTP81
#$$(info ==--==$(1),$$($1))
$(1):
#	echo ==--1:$(1)
#	echo ==--2:$$($$($(1)))
	$$($$($(1)))

endef

define testTP89
#$$(if $$($$($(1))),$$(info ==$$($$($(1)))))
$$(if $$($$($(1))),$$(eval $$(call testTP81,$1)))

endef

define testTP9
$$(eval $$(call testTP1,$(1)))
$$(eval $$(call testTP2,$(1)))

#$$(testList): 
#	echo $$($$($$@))
#	$$($$($$@))

endef

$(foreach aa1,$(testListEnv),$(eval $(call testTP9,$(aa1))))
showRunHelpListLast += $(testList2)
$(foreach aa1,$(testList2),$(eval $(call testTP89,$(aa1))))

ifdef dddddebug01
$(info aaaf1:$(aaaf1))
$(info aaaf2:$(aaaf2))
$(info aaaf3:$(aaaf3))
$(info bbbf1:$(bbbf1))
$(info bbbf2:$(bbbf2))
$(info bbbf3:$(bbbf3))
$(info rf1:$(rf1):$($(rf1)))
$(info rf2:$(rf2):$($(rf2)))
$(info rf3:$(rf3):$($(rf3)))
#$(info udp___base:$(udp___base):$($(udp___base)))
endif


ln:
	cd src2/ && rm -f base_*.go && ln -s ../src0/base_*.go ./
	cd src3/ && rm -f base_*.go && ln -s ../src0/base_*.go ./
	cd src4/ && rm -f base_*.go && ln -s ../src0/base_*.go ./

