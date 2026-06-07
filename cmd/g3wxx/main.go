// Copyright (c) 2026 Michael D Henderson. All rights reserved.

// Package main implements a command line tool to transform an original
// ASCII art Olympia map into a Worldographer data file that can be
// consumed by mapgen or woly to initialize an Olympia store.
package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"

	"github.com/mdhender/ottomap"
	"github.com/mdhender/ottomap/hex"
	"github.com/mdhender/ottomap/wog"
)

type City struct {
	Name      string // if blank, randomly generated
	Major     bool   // major or minor city
	SafeHaven bool
	Feature   string // filled in by code
}

var (
	sourceMaps = [][]string{
		{ // v1
			`................................................................................`,
			`................................................................................`,
			`...........p.mm..........pp...mmm...............................................`,
			`..........pppmm.........ssppmmmsmm..............................p...............`,
			`........pmmpppp.ffs....msppppmms......m...ff.....ffm...........ppppp............`,
			`.........mmmssppfpsf.....pppmmmm....ddmmfffffff.sfff..........msffpp...p........`,
			`.........mmpspppfffffm...ppmfm%...ddddfppppppfffffff1........mssfffpp.ppp.......`,
			`.........mmsspppsffffmm...fffm...dddddssppppmfffffpp........pspppffppppppp......`,
			`.....msspmmmsmpppffffm.....ff...mdddddsppppmmfffffpp........sspppfpppsppp.......`,
			`.....mssmppmmmppppmff.....mff...mddddddppppsmsmfssppp.......pppppfpppsppssf.....`,
			`.....mmmmfffppppsfffmf...mmfsf....dddddpppppmmmmmsppp.........ppfpppppppsfff%...`,
			`.......mmfffsssspfffff....ffff....dsdfdppppmmmsmmmssmmpff....pppfffpppppffsf....`,
			`......mmffsmmpssffpf4...........ddddffdpppmmfpssmmpmmmfff....pppffpspppfffsf....`,
			`......sffffmmmmsffpff...........dddmddppppmffmmsfmssfmfd.......pppfsppppfss.....`,
			`.......sffffmmmpfffsss.........dddddddppppmffmmmfffddmpdd.......pffsmmspmmpm....`,
			`....mmmssssmspppffspms..........dddddpppppmffmmmmmsdpppss.......ffmmmmsmmm......`,
			`....mmmsssmmsppppp......p...p....ddddppffffffffsffsdffps.......fffmmmfffsmm.....`,
			`....m.mmmssmssppp......pppp.ppp...fdddffppmmffsssssddfffss.....ffmm.sfffss......`,
			`......mm.smmsppp.......pppppppp....ddffpppmmffmspddddffffp......p....fff........`,
			`...........mspp.......ppppppssspp...dfffpmmmfffssddddddfmmmm.........fs.........`,
			`...........sspp.......pppffpmmppp....pfffffffffmsssdddddddmm.........fss........`,
			`............s.......pppfffffmmppp.....pfffffffmmmfsddddpdddm...f.pp...m.........`,
			`....................pfpfsfffmmpppff....mfpppffmmmfffdppppddd...mpspp............`,
			`....................pfffffffffpfpf....mmfdpfffpppffsspppppd...mmppppp...........`,
			`............dddd....fffffdffffsffff...mmpppfpppsdfffppppppp...mmmfffpp.p........`,
			`...........ddddfff...fdpddffffffff%...mmppppppmfffffpp.p.....fmmfffffffppp......`,
			`.......dd..dddmfff...fdddddfpfppff...fmmmmppppmmfpfmpm......fmmppppffffffpp.....`,
			`......dddssdddmmmff....dddd.ppp.....mddmppppppmmffmmmmm....fffpppppffffff.......`,
			`......ddddddfmmpmfff...dddd........mssfffpsfppmmmmmmm....pfffpppmppffffff.......`,
			`.....mmmdddfffmfff%...pdddd.......mmmfffffffpmmm........spfffppfmsssffffpp......`,
			`......mmmddpfmffmfff.....d.......mmmm.ff%.f............fssfffppmmmsssffff.......`,
			`......ffmfdffmfffffff..............m...f................sppffppppmpppp.f........`,
			`....f.fffddffmpffsfffp......pppm................pp........ppmsppmmmpp...........`,
			`...ffffpfppffmmffsffffss....mppp.............ff.ps.........mmspppsmp.....s.s....`,
			`...mmmpppppppmfffpppffs.....mpppppp..p......ffpppss..p%p....mmmppff......pps....`,
			`..fmmmmpppsspffffppppp........ppspppppp.....fppppppp..pppp...pppp6...ff.ppppp...`,
			`..mmmmpfpppssssfppppp............pffff.......ppppppffssppp...pppff...fffppp.....`,
			`..mmmmpppmppsfpfsssp...ppp.......ppmf%.......pppppfffsppp.....pppp...ffpppps....`,
			`..mppppppmmmfffffpp...ppp.........pff.......ppppppffffppp.....ps....pfffppf.....`,
			`..ppsspppmmmffsfss...pppp.pmmfff...f.........ppppffffsssp......s...ppsfppffpp...`,
			`...ppppppmmsffsffs...ppppppfffff......dd.....ppppffddpspp.........mpppppfffppp..`,
			`..ppppppppmppffff....ppsppffffffff...ddd....ppmmpfdddpsp..........mpppppffpp....`,
			`..ppppdpfppppfss....mppsppff%..ff...dddddd...pmOmffdpsspp........mpppmmfffmpp...`,
			`.....ddffspppffff....mmppffff..f...dddddddd...mmmmddd..........sfmppppmmmmmssp..`,
			`.....sddddddpfff3...mmmmmfffff....ffffddddd...ppmdddddddp....mfffffpppmppmpfpp..`,
			`......ddddddppff....mssmmffff.....%fffddppm...p.ppd.ddd.........fffppppppmpf....`,
			`.....pdddppppppf...pssmffffff...s.fffmddppmm........d.....mm......fppppppmfff...`,
			`....ppppdpppppppp...p.mfpff....ssmffmppppmmmp............pmmm....mfffppppmff%...`,
			`......pfppppppm.........pp.......mmmsppmmmmmpp.........pffpmmmp....fpp.pfffff...`,
			`......pfffppffff...............sssmmsppmmmmmppms.....fffffpsspff....p.....m.....`,
			`.....ffffmmpffff..................mmpppmfffmpffs.....pffpppsffff...ss...........`,
			`.......fffmm..ff...................ppppfffffpf........ffppsp.%ff............p...`,
			`.........sm...ff.............mf....ppppppmmfff........ffffss..ff......ffm..ppp..`,
			`..........m....m......fmm.ssfmfs...pppppmmffss.......................fffmspppp..`,
			`....smm...m...........pppmmmfffs...pp.ppmmmmms......m.................%fmmpspp..`,
			`...msppm.............sfmdmmmfffss......smmmm.......mm...........fff....fsmmppp..`,
			`..ppsppms...m......%fffmddmmfffff.......pm..........msf..ssmfffmsfff....sfmmp...`,
			`..ppppmmsf.fff....ffffffddmmmffffp......p...........mss..fsffffssff%.....ffppf..`,
			`..pppppssffffff...fpfffsdddddffff...........f%...fffmssm.mmmmffppffffpp...fspf..`,
			`..ppppfffsff.....ffpppfsdddddpfff.........sfff...ffsssmmmmmmmmpppppfpp.....sp...`,
			`...pppfmmmf%.....5fpppfsdddddppfpf.......psfpfs....fmmmmmmmpppppppppp......fff..`,
			`...pppfmmmff.....ffppffpsdsssppfpp.....pppffffs...sssmsffmfffppfffffp......f....`,
			`.....fffmm........ppfffppssmfpppp........ffffpss...mmmfffmfssmmffffppp..........`,
			`.......fm.........ppfmffpmmfffpfp.........fppp....smmmfmffffmmfffffspp..........`,
			`...................ppffpppmmffsf...........ppp...fffmmmmmmfmmfffsmsssps.........`,
			`....f.....f.fff....ppffpppmmmfff....m.ppp...pp...2ffffmmmsfmmffffmmms...........`,
			`....ffm..pfffffs...pffffffpmmsf....mmsppp....p...ffffdmmmsffmfssmmmmsm..........`,
			`....ffpppppfffff.....ffffpp..sm...sppspppdd........dfdsdfsffffffmsmms...mmm.....`,
			`....ffpppppffff......ffpss...s......ppppsdd.......ddddddfffffffmmmmmf...smmm....`,
			`...mmmmmppsffff...................sppppppd.......ssddddddfffffffmmmmm...mmmm....`,
			`.....fmsppfffff....................ffffpddd......sfffddddsfddfsssmssm...mmmmpp..`,
			`......fppppff......mf.............mfffmmdddf.....ffffddsdpdddfsmssssm....mmppp..`,
			`....%fpppppff....ffmfff...........mf%ffddds....ssffffddffppddmmmffss.....ppppp..`,
			`.....fppppffff...ffffffpppp........f.ffddd.........ffddffppdmmmmffff...fpppp....`,
			`....ffpffff.....%fffffppp.............fssd........mffsfffppddsmmf.f...ffffppp...`,
			`....f.fffff.....ffppppp................ss..........ffppppmpddsff.....f%fffppp...`,
			`.................mppppp.............................ppppppppdsffff......spppp...`,
			`.................mppppp............................ppppppppddffff........ff.....`,
			`................................................................................`,
			`................................................................................`,
		},
		{ // v2
			`............................................                                   `,
			`............................................                                   `,
			`.............................^^^............ ooooo&&             ooooooo.......`,
			`,...........ooo^............^^*^...........   ooo^^o&&&&       ooo^^^^ooo.sm..,`,
			`,,,,ooo.....^oo^^...........^^^^...........    oooo^^oo&&&     ooo^^^oooo..^,,,`,
			`,,,ooooo.....^^^*............^^.........oooo    oooo^^^#^&&     oo^^#^oo...&&,,`,
			`,,,oo^^ooo.............................o^?&&&    ooooo^^?&&     &&&?&&&&....*,,`,
			`,,,oo?^^oo............^^^............. oo^^&&&    ooooo^^o&&     &&&&&&&...&&,,`,
			`,,,,oo^^^ooo.........oooo............   oo^#?&&    oooo^^o&&      &&&&?&&...^,,`,
			`,,,,,ooo^^ooo^......oo^^oo.......^^^^    oo#^&^*    ooo^^o&&&     &&&&^&*,,,,,,`,
			`,,,,,,,oooooooooooooooooooo....ooooooo    ooo^&&&    o^?^^&o*      &&&&&&&,,,,,`,
			`,,,,,,,oooooo^^ooooooo^^ooooo^?^o^ooooo    oo&&&&    oo^^o&&       &&&?ooo,,,,,`,
			`,,,*,,,^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^      &&&     ooo&&         &&&&oo,,,,,`,
			`,,^^^,,#####?###########?#############^^oo                   oooo         ,,,,,`,
			`,,,^o,,^^^^^^^^^^^^^^^^^^^^^^^^^^^?^^^^oooo&&     o*&       oooooo         ,,,,`,
			`,,,o^,,^ooo^oo^ooo^ooo^^^^^^^^&o&&&&&?&ooo&oo&    &&&&    &ooo^^oo&        ,,,,`,
			`,,,,,,,,ooooooooooooooooo^oo~~~~o&&&&&&ooooo&&    &&?&    &&&?^&&&&      oo,,,,`,
			`,,,,,,,,~~~~~~~~~~oo^oo^oo~~~~~~~~&&%%%%ooooo&    &&&&     o*&&&&&      *oo^^,,`,
			`,ss,,,oooooo~~^~~~&&&ooo&~~~~&&~~~~%%%%oooo&&&^    &&&                 ooo^ooo,`,
			`s,,,,oooooooo~~~~&&?&&o&~~~~&&&~~~~~oooooo&?&&&                        o^^ooooo`,
			`,,,,,&&&&&&&oooo&&?&oo&&~~~&&?&~~~~~oooooo&&&&&                       &o^?ooooo`,
			`,,,,,&&&&&&?&ooo^^^oo&&&~~~&&&&~~~~~&&ooooo&&&                &&       &&?oo&&,`,
			`,,,,,,,&&&&&&&ooooooo&&~~~~&#&&~~~~~&&&ooo&&&                 oo&&       &&&&&,`,
			`,,,*,,,~~~~~~~oooooo&&&~~~&&&&~~~~~~~&&&&&&&                 ooooo&~~~    ,,,,,`,
			`,&&&&,,~~~^~~~oo^^&oo&&~~~~*&~~~~~~~~~&&&&                  oo^^ooo~~~~~  ,,,,,`,
			`,&&&,,,~~~~~~ooo#&&oo&&~~~&&&~&&&~~~~~!!!!       &&&        ooo^^oo~~~~~~~,,,,,`,
			`,,,,,,&&oooooooo?oooooo~~~&&~~~*&~~~~&&!!&&    .&&&o&       *oooooo~~~~~~~&&,,,`,
			`,,,,,&&&oooooo~~&oooo&&~~~~~~~&&&~~~&&&oo&&   ..&oooo&&    ooooooo~~~~~~~&&&&,,`,
			`,,,,&&!!#&&oo~~~~&&&&&&&~~~~~~~~~~~~&&&oo&&......ooo&oo&o^^oooooo~~~~~*~~~&&&,,`,
			`,,,&&!!&?&&&~~~~~&&?&&&&&~~~~~~~~~~~~&&&&&&.......oooooooo^^#?ooo~~~~^^~~~&&&,,`,
			`,,,&&&!!&&&&~~^d~&&&&&&&&~~~~~^s~~~~~&&&&&   .......ooooooo^^??oooo~~~~~&&&&,,,`,
			`,,,,&&&&&&&&~~~~~~&&&&&&~~~*s~~~~~~~~&&o&&    ......oooooooo^^?ooooo~~~~oo&,,,,`,
			`,,,,,&&&&&&         ~*&               &&&      ........oooooo^^^oooooooooo&,,,,`,
			`,                                               ........&&&&oo^^^ooooooo&&&...,`,
			`   d*         %%%%%                              .........&&&&oo^^^^ooo&&&&&...`,
			`             %%??%%                               .ooo.....&&&&o^##^^o&&?&&&...`,
			`             %%%%%%        ooo                    .&&oo....&&&oo^^#^^o&&&&&&...`,
			`  o*          %*%?%%%     &&oooooooo              ..oo*...&&?&oo^^?^^o&&&?&&...`,
			`     &&&&      %%%%%%     &&&&oooo*o&          ooo.........&&&&o^^#^^^o&&&&&...`,
			`    oo#&&&                 &&&&&&oooo       ~~~o#oo........&&&&oo^^#^^o&&&&....`,
			`   &^^o&o*                  &&&&&&&&   oooo~~~~oooo........&&&oooo^#^^ooooo....`,
			`   ?&&&&&                              o#ooo~~ooooo......&&&&&oooo^#^^oooooo...`,
			`                                        ooooooooooo.......&&&&!!!o^#^ooooo&&...`,
			`    !!!        ooooo                     oooooo*oo.......&&&&&!?!o^?^oo!!!&&...`,
			`   ooooo     ooooooooo        ooooooo      oooooo........&&&&&!!!o^#^ooo!!!&...`,
			`   ^^^^     &&&oooooooooo!!!!oooooooooo     ...........~~~&&&&&&o^^^^ooooo&....`,
			`  %%%%     &&?&ooo^^ooooo!!!ooooo^?oo&&&   .........~~~~~~~~&&&o^^?^^^ooo&.....`,
			`  %%^*    &&?&&o^^^^^oo*o!!ooooo^^^oo&&&  .....&&&~~~~~~~~~~~&&&o^^^oooo&~~~...`,
			`           &&&oo^^#^^oooo!!!o&&^^^oo&&&   ...&&&&&&~~~~~~~~~~&&&oo^^oooo&~~^^..`,
			`            &&&&&oooooo&&&!!?&&&&&&&&&~~~~..&&&^?&&~~~~~~~~&&&ooo^^oooo&~~^^%..`,
			`             &&&&&&oo&&&&~~~~&&&&&&~~~~~~~..&&&??&&&~~~~~&&&oo?^^^^ooo&~~^^%...`,
			`               &&&&&&&&~~~~~~~~~~~~~~~%%%~...&&&&&&&~~~&&&oo^^^^oooo&&~~^^%....`,
			`   *^^^^        ~~~~~~~~~~~~~~~~~~~~~~~~*~....&&&&&%%%o&ooo^^^ooo&&&~~~%^%.....`,
			`  ^^^^^^o      ~~&&&~~~~~~~~~~~~~~~~~~~~~~~....&&&%%%%oooo^^ooo&&&~ms~%^^......`,
			`  ^^??^^^     ~~~&&oo~~~~*~~~~~~~&&&oo~~~~~~.....&%%%%%%oooo%%%...~~~%%^.......`,
			`~~^^??^^^    ~~~~~oo*~~~%%~~~~~~&&&ooo~~~~~~~......%%%%?%%%%%........*%......~~`,
			`~~~^^^^^~~~~~~~~~~~~~~~~~~~~~~~~~~*&&&~~~~~~~~......%%%%%%%...............&&&~~`,
			`~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~..............            &.*&&~`,
			`~~~~~~~~~~~&&&&&&&&~~~~~&&&&&&&&&&&~~~~~~~~~~~~~..........               &&&&&~`,
			`~~~~~~~&&&&&&&&&&&&ooooo&&&&&&&?&&&&&&&~~~~~~~~~~                         &&&~~`,
			`~~~~~&&&o?^^^^^^ooo&&&&&&oooo&&?&&&&&?&&~~~~~~~~~                        ~~~~~~`,
			`~~~~&oo^^^^^*^oooooooooooooo&&?&&&ooo&&&~~~~~~~~~                 oooooo~~~~~~~`,
			`~~~~&oo^^oo^^^oo&&&&&&&&&&&&&&&&ooooo&&&~~~&&~~~~              ooooooo&ooo~~~~~`,
			`~~~&oo^?^o^^^#^&&~~~~~~~~~~~~~&&&&oooo&&&~~&*~~~~             oooooo&oooooo~~~~`,
			`~~~&oo^^^^^^^^&~~~~~~~~~~~~~~~~&&&&oooo&&~~~&&~~~            *ooo&oooooooooo~~~`,
			`~~~&oo^^^^^o&&~~~~~&&&&*~~~~~~~~&&&oooo#&&~~~~~~~           oooooooooo&ooooo~~~`,
			`~~~~&oo&&&oo&~~~~&&&&&oo&&&&~~~~~&&&&&&?&&~~~~~~~          oooo?oo&o%%%%%%oo~~~`,
			`~~~~&&&&?&&&~~~~&&&&ooo^oo&&&~~~~~&&!!&&&&~~~&&~~~         oooooooo%%%?%#?&&o~~`,
			`~~~~~&&&&&~~~~~&&&&&o^^^ooo&&~~~~~&&&&&&&~~~&&&&~~         oo&oooooo%%%%%%&&&~~`,
			`~~~~~~~~~~~~~~~&&&?&o^#&oo&&~~~~~~~&&&&&~~~&&&?&~~          ooooooo&o&&&&&?&&~~`,
			`~~~~~~~~~~~^^~~~&&&&oo&?&&&~~&&~~~~~~~~~~~&&&#&&~~          ooooooooo&?&&&&&~~~`,
			`~~~^^~~~~~~~~~~~~&&&&&&&&&&~&*~~~~~~~~~~~~~~*&&~~~           ooo&oooo&&&&&&&~~~`,
			`....*~~~~~~~~~~~~~~&&&&&&*~~&&~~~~~&&&&~~~~~~~~~~~            oooooooo&&&&&&~~~`,
			`..........~~~~~~~~~~~~~~~~~~~~~~~~&&&&&&........................ooo&ooo&&&&....`,
			`..oooo......^ooo^...................*&&................%%%.......ooooo&&&&.....`,
			`.oooooo...^oooooooo..........................oo^......%%%%%....................`,
			`oooo*oo..ooooooooooooooooo^................^^oooo...o^%%%%%%...................`,
			`oooooo...ooo&&&&&&&&&&&oooooo.............^oooooooooooo%%%%ooo.................`,
			`.oooo...oo&&&&^^^^^^^^&&&oooooooooooo^^~~^ooooooo^oooooooooooo^^^^oo^..........`,
			`..oo...oo&&&^^^^^^^?#^^&&&&&&&&&&oo^ooo~~oooooooooooooooooooo^^^^oooooo^.......`,
			`......^oo&&^^^^?^^^#^^^^^^^&&&^^&oooooo~~oooooo^oooo^ooo^oooo^?^ooooooooo^.....`,
			`.....^o&o&&^^^^^?##^^?^^^^^^^^^&&ooo&&&~~&&&&ooooooooooooooo^^^^ooooo^ooooo....`,
			`....^ooooo&&^?^#ooo#^^^^?^^^^&&&oooo&&~s~~&&&&oooooooo^ooo^^^^oooooooooooo^....`,
			`....^ooooo&&^^^#ooo#^?^^^^&&&&&ooooo&&~~&&~~~&&&&&&&&&oooo^^^#^ooo^ooooo*ooo...`,
			`....^o&oo&&^^?^^##?^^^^^&&&&&&oo^oo&&&~&&&&&~~~~~~~~&&&ooo^^^^^ooooooooooooo^..`,
			`....^oooo&&^^^^?^^^^^?^&&&&&ooooooo&&~~&&&&&&&&&&&&~~&&&&&&oo^^ooooo&&&&&&&^^..`,
			`.....^oooo&&&^^^^^^^^^&&&&oooo*ooo&&~~&&&&&&&&&&?&&&~~~~~~~&&^ooooo&&&&&&&&&...`,
			`......^oooo&&&&&&&&&&&&&&oooooo&&&&~~&&&&&??&&&&&&&&&&&&&&~~&^oooo&&?&&&&&&....`,
			`.......^oooo&&&&&&&&&&&oooo&oo&&&~~~&&&?&&&&&&?&#&&&&&&?&&&~~^^ooo&&&&!!!!!....`,
			`..&&....^oooooooooooooooooooooo&&~~&&&&&&&#&&&#&&&&#&&&&&?&~~^^ooo&&&!!!!!!....`,
			`.&&^^....^^ooo^oooooo^^o&ooo^oo&~~&&&&&&&&&&&&&&&&&&&&?&&&&~~~~^^^&&&!!!!!.....`,
			`.oo^^^...........ooooo^ooo^oooo...&&&&&&&?&&&&&&&&&&&&&&&&........^&&!!!.......`,
			`.ooo^^*................................&&&&&&&&*...............................`,
			`..oo^^^........................................................................`,
		},
		{ // v3
			`....................................................................................................`,
			`...........oooooooooooooooo...................................ooooooooooooooooo.....................`,
			`........oooooooooooooooooooooo...........oooooooooo.......ooooooooooooooooooooooo.....oooo..........`,
			`......ooooooooooooooooooooooooo........oooooooooooo......ooooooooooooooooooooooooooooooooooo........`,
			`.....oooooooooooooooooooooooooo........oooooooooooooo....ooooooooooooooooooooooooooooooooooooo......`,
			`...ooooooooooooooooooooooooooo..........ooooooooooooooo...oooooooooooooo.....oooooooooooooooooo.....`,
			`...oooooooooooooooooooooooooo............ooooooooooooooooooooooooooooooo.....ooooooooooooooooooo....`,
			`..ooooooooooooooooooooooo...................oooooooooooooooooooooooooooooo....ooooooooooooooooooo...`,
			`..ooooooooooooooooooooo.........ooo.........oooooooooooooooooooooooooooooo.....oooooooooooooooooo...`,
			`..ooooooooooooooooooo........ooooooo........oooooooooooooooooooooooooooooooo.....ooooooooooooooooo..`,
			`..ooooooooooooooo.........oooooooooo........oooooooooooooooooooooooooooooooooo.....ooooooooooooooo..`,
			`..ooooooooooooo.......oooooooooooooooooo.....ooooooooooooooooooooooooooooooooo.....oooooooooooooooo.`,
			`...oooooooo.......oooooooooooooooooooooo......oooooooooooooooooooooooooooooooo....ooooooooooooooooo.`,
			`....oooo.......ooooooooooooooooooooooooooo......oooooooooooooooooooooooo..........ooooooooooooooooo.`,
			`............ooooooooooooooooooooooooooooooo...............oooooooooooo.............oooooooooooooooo.`,
			`.........ooooooooooooooooooooooooooooooooooo.......................................oooooooooooooooo.`,
			`.....ooooooooooooooooooooooooooooooooooooooo.....ooooooo.............................ooooooooooooo..`,
			`...oooooooooooooooooooooooooooooooooooooooo....oooooooooooooooooooooooooo.............ooooooooooo...`,
			`...oooooooooooooooooooooooooooooooooooooooo....oooooooooooooooooooooooooo.....................ooo...`,
			`...oooooooooooooooooooooooooooooooooooooooo....oooooooooooooooooooooooooooooooooooooo...............`,
			`...oooooooooooooooooooooooooooooooooooooooo....ooooooooooooooooooooooooooooooooooooooooo............`,
			`...oooooooooooooooooooooooooooooooooooooo.....oooooooooooooooooooooooooooooooooooooooooooooo........`,
			`...ooooooooo.....................oooooooo....ooooooooooooooooooooooooooooooooooooooooooooooooo......`,
			`...ooo..............................ooo.....ooooooooooooooooooooooooooooooooooooooooooooooooooo.....`,
			`.............oooooooooooooooooo............ooooooooooooooooooooooooooooooooooooooooooooooooooooo....`,
			`........oooooooooooooooooooooooooo.......oooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`......oooooooooooooooooooooooooooooo.....oooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`.....oooooooooooooooooooooooooooooooo.....ooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`.....oooooooooooooooooooooooooooooooo.....oooooooooo.............oooooooooooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooooo.....ooooooo..................ooooooooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooooo...................................oooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooo.......................................oooooooooooooooooooooo...`,
			`.....ooooooooooooooooooooooooooooo...........................................ooooooooooooooooooo....`,
			`.....ooooooooooooooooooooooooooo.........ooooooooo.........oooooooooooo.........ooooooooooooooo.....`,
			`.....ooooooooooooooooooooooooo........oooooooooooooo.....ooooooooooooooo..........oooooooooooo......`,
			`.....oooooooooooooooooooooo.........ooooooooooooooooooooooooooooooooooooo...........................`,
			`.......oooooooooooooooooo.........oooooooooooooooooooooooooooooooooooooooooo........................`,
			`........oooooooooooooooo.......oooooooooooooooooooooooooooooooooooooooooooooo.......................`,
			`..........ooooooooooo........oooooooooooooooooooooooooooooooooooooooooooooooo.......oooooo..........`,
			`..oooo......................ooooooooooooooooooooooooooooooooooooooooooooooopo.......ooooooooo.......`,
			`.ooooooo...................ooooooooooooooooooooooooooooooooooooooooooooooopf.......ooooooooooooo....`,
			`.oooooooooo..............fooooooooooooooooooooooooooooooooooooooooooooooofff.......oooooooooooooo...`,
			`..oooooooooooo.........1pfpoooooooooooooooooooooooooooooooooooooooooooooppp2.......ooooooooooooooo..`,
			`...ooooooooooo........pffpoooooooooooooooooooooooooooooooooooooooooooooooffp........oooooooooooooo..`,
			`...oooooooooo........offpfoooooooooooooooooooooooooooooooooooooooooooooooofpf........ooooooooooooo..`,
			`...ooooooooo........ooofooooooooooooooooooooooooooooooooooooooooooooooooooopoo.......ooooooooooooo..`,
			`...ooooooo........ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`...oooooo.......ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`....oooo.......oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`.............oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo.......oooooooooooo..`,
			`...........ooooooooooooooooooooooooooooooooooooooopppoooooooooooooooooooooooooo........oooooooooo...`,
			`...........ooooooooooooooooooooooooooooooooooooppppppffoooooooooooooooooooooooo........ooooooooo....`,
			`..........oooooooooooooooooooooooooooooooooooooppppppfffoooooooooooooooooooooo..........ooooooo.....`,
			`........ooooooooooooooooooooooooooooooooooooooopppO4ffffoooooooooooooooooooooo......................`,
			`........oooooooooooooooooooooooooooooooooooooooppppppfffooooooooooooooooooooooooo...................`,
			`.......ooooooooooooooooooooooooooooooooooooooooppppppffooooooooooooooooooooooooooooooo..............`,
			`.......ooooooooooooooooooooooooooooooooooooooooooopppoooooooooooooooooooooooooooooooooooo...........`,
			`........ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo...........`,
			`.........oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo...........`,
			`..........fooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo............`,
			`..........fpoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo............`,
			`..........fppoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo.............`,
			`..........5fppooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo.............`,
			`...........ffooooooooooo.......ooooooooooooooooooooooooooooooooooooooooooooooooooooooo..............`,
			`............oooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooooooooooo...............`,
			`.............ooooooooo............oooooooooooooooooooooooooooooooooooooooooooooooo..................`,
			`..............ooooooo.............oooooooooooooooofoooooooooooooooooooooooooooo.....................`,
			`................ooooo..............oooooooooooooopppooooooooooooooooooooooooo.......................`,
			`........................................oooooooofpppfooooooooooooooooooooo..........................`,
			`............................................ooofff3fffoooooooooooooooooo.........oooooooooooo.......`,
			`...............................................................................oooooooooooooooo.....`,
			`.........oooooooooo........................................................oooooooooooooooooooooo...`,
			`........oooooooooooo..................................................oooooooooooooooooooooooooooo..`,
			`......oooooooooooooooo.............................................ooooooooooooooooooooooooooooooo..`,
			`.....oooooooooooooooooo..........................................oooooooooooooooooooooooooooooooooo.`,
			`....oooooooooooooooooooo.........ooooo..........................ooooooooooooooooooooooooooooooooooo.`,
			`...ooooooooooooooooooooo........oooooooo.......................oooooooooooooooooooooooooooooooooooo.`,
			`...ooooooooooooooooooooo.......ooooooooooo....................ooooooooooooooo........oooooooooooooo.`,
			`...ooooooooooooooooooooo.......ooooooooooooooooooooooooooooooooooooooooooo............ooooooooooooo.`,
			`...oooooooooooooooooooo.........ooooooooooooooooooooooooooooooooooooo%.................oooooooooooo.`,
			`.....ooooooooooooooooo...........oooooooooooooooooooooooooooooooooooo..................ooooooooooo..`,
			`......oooooooooooooo..............oooooooooooooooooooooooooooooooooo..................oooooooooooo..`,
			`...................................ooooooooooooooooooooooooooooooo.......%ooooo.......oooooooooooo..`,
			`.....................................oooooooooooooooooooooooooooo......oooooooo.......ooooooooooo...`,
			`..........oooo...oooooooooooo........oooooooooooooooooooooooooo.......oooooooo........oooooooooo....`,
			`.......oooooooooooooooooooooooo.......ooooooooooooooooooo.............ooooooo.........oooooooooo....`,
			`......oooooooooooooooooooooooooo.......ooooooooooooooooo...............ooooo.........oooooooooo.....`,
			`....ooooooooooooooooooooooooooooo.......ooooooooooooooooooooo......................ooooooooooo......`,
			`....ooooooooooooooooooooooooooooo.......oooooooooooooooooooooo....................oooooooooo........`,
			`......ooooooooooooooooooooooooooo.......oooooooooooooooooooooooo.................oooooooooo.........`,
			`.......oooooooooooooooooooooooooo.......oooooooooooooooooooooooooo.............ooooooooooo..........`,
			`.......ooooooooooooooooooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`......oooooooooooooooooooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`......ooooooooooooooooooooooooo..........oooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`....oooooooooooooooooooooooooo..........oooooooooooooooooooooooooooooooooooooooooooooooooo..........`,
			`....ooooooooooooooooooooooooo...........ooooooooooooooooooooooooooooooooooooooooooooooooo...........`,
			`....ooooooooooooooooooooooo.............oooooooooooooooooooooooooooooooooooooooooooooooo............`,
			`....ooooooo...ooooooooooooo..............oooooooooooooooooooooooo...............ooooooo.............`,
			`.....ooooo..........................................................................................`,
			`....................................................................................................`,
		},
		{ // v4
			`....................................................................................................`,
			`...........oooooooooooooooo...................................ooooooooooooooooo.....................`,
			`........oooooooooooooooooooooo...........oooooooooo.......ooooooooooooooooooooooo.....oooo..........`,
			`......ooooooooooooooooooooooooo........oooooooooooo......ooooooooooooooooooooooooooooooooooo........`,
			`.....ooooooooooooooooooooooooo%:::.....ooooooooooooo%....oooooooooooooooooo%oooooooooooooooooo......`,
			`...ooooooooooooooooooooooooooo...::.....ooooooooooooooo...oooooooooooooo.....oooooooooooooooooo.....`,
			`...oooooooooooooooooooooooooo.....::::...ooooooooooooooooooooooooooooooo.....ooooooooooooooooooo....`,
			`..ooooooooooooooooooooooo............::.....oooooooooooooooooooooooooooooo....ooooooooooooooooooo...`,
			`..ooooooooooooooooooooo.........ooo...:::...oooooooooooooooooooooooooooooo.....oooooooooooooooooo...`,
			`..ooooooooooooooooooo........ooooooo....::..oooooooooooooooooooooooooooooooo.....ooooooooooooooooo..`,
			`..ooooooooooooooo.........oooooooooo.....::.oooooooooooooooooooooooooooooooooo.....ooooooooooooooo..`,
			`..ooooooooooooo.......oooooooooooooooooo..:..ooooooooooooooooooooooooooooooooo.....oooooooooooooooo.`,
			`...oooooooo.......oooooooooooooooooooooo..:::.oooooooooooooooooooooooooooooooo....ooooooooooooooooo.`,
			`....oooo.......ooooooooooooooooooooooooooo..:...oooooooooooooooooooooooo..........ooooooooooooooooo.`,
			`............ooooooooooooooooooooooooooooooo.::............oooooooooooo.............oooooooooooooooo.`,
			`.........ooooooooooooooooooooooooooooooooooo.:.....................................oooooooooooooooo.`,
			`.....%oooooooooooooooooooooooooooooooooooooo.:...ooooooo.............................ooooooooooooo..`,
			`...oooooooooooooooooooooooooooooooooooooooo..:.oooooooooooooooooooooooooo.............ooooooooooo...`,
			`...oooooooooooooooooooooooooooooooooooooooo..:.oooooooooooooooooooooooooo.....................ooo...`,
			`...oooooooooooooooooooooooooooooooooooooooo..:.oooooooooooooooooooooooooooooooooooooo...............`,
			`...oooooooooooooooooooooooooooooooooooooooo.::.ooooooooooooooooooooooooooooooooooooooooo............`,
			`...oooooooooooooooooooooooooooooooooooooo..::.oooooooooooooooooooooooooooooooooooooooooooooo........`,
			`...ooooooooo.....................oooooooo.::.ooooooooooooooooooooooooooooooooooooooooooooooooo......`,
			`...ooo..............................ooo..::.ooooooooooooooooooooooooooooooooooooooooooooooooooo.....`,
			`.............oooooooooooooooooo........:::.ooooooooooooooooooooooooooooooooooooooooooooooooooooo....`,
			`........oooooooooooooooooooooooooo.....:.oooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`......oooooooooooooooooooooooooooooo...:.oooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`.....oooooooooooooooooooooooooooooooo..:..ooooooooooooooooooooooooooooooooooooooooooooooooooooooo...`,
			`.....oooooooooooooooooooooooooooooooo..:..ooooooooo%:............oooooooooooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooooo.:...ooooooo..::..............ooooooooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooooo.:.............:...................oooooooooooooooooooooooo...`,
			`....oooooooooooooooooooooooooooooooo..::...........:::.....................oooooooooooooooooooooo...`,
			`.....ooooooooooooooooooooooooooooo....:............:.........................%oooooooooooooooooo....`,
			`.....ooooooooooooooooooooooooooo......:..ooooooofp.:.......oooooooooooo......:..ooooooooooooooo.....`,
			`.....ooooooooooooooooooooooooo........%oooooooofpff6.....ooooooooooooooo.....::...oooooooooooo......`,
			`.....oooooooooooooooooooooo.........ooooooooooofpfpfpfooooooooooooooooooo.....::....................`,
			`.......oooooooooooooooooo.........oooooooooooooofffffffooooooooooooooooooooo...:....................`,
			`........oooooooooooooooo.......oooooooooooooooooooooooooooooooooooooooooooooo..:....................`,
			`..........oooooo%oooo........oooooooooooooooooooooooooooooooooooooooooooooooo..:....oooooo..........`,
			`..oooo..........:...........ooooooooooooooooooooooooooooooooooooooooooooooopo..:....ooooooooo.......`,
			`.ooooooo........:::........ooooooooooooooooooooooooooooooooooooooooooooooopf..::...ooooooooooooo....`,
			`.oooooooooo.......:......fooooooooooooooooooooooooooooooooooooooooooooooofff..:....oooooooooooooo...`,
			`..oooooooooooo....:::::1pfpoooooooooooooooooooooooooooooooooooooooooooooppp2.::....ooooooooooooooo..`,
			`...ooooooooooo........pffpoooooooooooooooooooooooooooooooooooooooooooooooffp........oooooooooooooo..`,
			`...oooooooooo........offpfoooooooooooooooooooooooooooooooooooooooooooooooofpf........ooooooooooooo..`,
			`...ooooooooo........ooofooooooooooooooooooooooooooooooooooooooooooooooooooopoo.......ooooooooooooo..`,
			`...ooooooo........ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`...oooooo.......ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`....oooo.......oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......ooooooooooooo..`,
			`.............oooooooooooooooooooooooooooooooooo8ooooooooooooooooooooooooooooooo.......oooooooooooo..`,
			`...........ooooooooooooooooooooooooooooooooooooooopppoooooooooooooooooooooooooo........oooooooooo...`,
			`...........ooooooooooooooooooooooooooooooooooooppppppffoooooooooooooooooooooooo........ooooooooo....`,
			`..........oooooooooooooooooooooooooooooooooooooppppppfffoooooooooooooooooooooo..........oooo%oo.....`,
			`........ooooooooooooooooooooooooooooooooooooooopppO4ffffoooooooooooooooooooooo..............:.......`,
			`........oooooooooooooooooooooooooooooooooooooooppppppfffooooooooooooooooooooooooo...........:.......`,
			`.......ooooooooooooooooooooooooooooooooooooooooppppppffooooooooooooooooooooooooooooooo......:.......`,
			`.......ooooooooooooooooooooooooooooooooooooooooooopppooooooooooooooooooooooooooooooooooo%::::.......`,
			`........ooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo...:.......`,
			`.........oooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo...:.......`,
			`..........fooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo....::......`,
			`..........fpoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo.....:......`,
			`..........fppoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......:......`,
			`..........5fppooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo......:......`,
			`..........:ffooooooooooo.......ooooooooooooooooooooooooooooooooooooooooooooooooooooooo.......:......`,
			`..........:.oooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooppppffooo........:......`,
			`..........:..ooooooooo............oooooooooooooooooooooooooooooooooooooooopppppfff...........:......`,
			`........:::...ooooooo.............oooooooooooooooofoooooooooooooooooooooopppff7..............:......`,
			`........:.......oooo%:.............oooooooooooooopppooooooooooooooooooooopppf.::::::::::::::::......`,
			`.......::............:..................oooooooofpppfooooooooooooooooooooo..................:.......`,
			`.......:.............::::...................ooofff3fffoooooooooooooooooo.........ooooooooooo%.......`,
			`.......:................::::......................:............................oooooooooooooooo.....`,
			`.......:.oooooooooo........::::...................:.....:::::::::..........oooooooooooooooooooooo...`,
			`.......:oooooooooooo..........:...................:....::.......::::..oooooooooooooooooooooooooooo..`,
			`......o%oooooooooooooo........::::...............:::::::...........%oooooooooooooooooooooooooooooo..`,
			`.....oooooooooooooooooo..........:..............::...............oooooooooooooooooooooooooooooooooo.`,
			`....oooooooooooooooooooo.........%oooo.........::...............ooooooooooooooooooooooooooooooooooo.`,
			`...ooooooooooooooooooooo........oooooooo......::...............oooooooooooooooooooooooooooooooooooo.`,
			`...ooooooooooooooooooooo.......ooooooooooo....:...............ooooooooooooooo........oooooooooooooo.`,
			`...ooooooooooooooooooooo.......ooooooooooooooo%ooooooooooooooooooooooooooo............ooooooooooooo.`,
			`...oooooooooooooooooooo.........ooooooooooooooooooooooooooooooooooooo%.................oooooooooooo.`,
			`.....ooooooooooooooooo...........oooooooooooooooooooooooooooooooooooo..................ooooooooooo..`,
			`......oooooooooooooo..............oooooooooooooooooooooooooooooooooo..................oooooooooooo..`,
			`...................................ooooooooooooooooooooooooooooooo.......%ooooo.......oooooooooooo..`,
			`.....................................oooooooooooooooooooooooooooo......oooooooo.......ooooooooooo...`,
			`..........oooo...oooooooooooo........oooooooooooooooooooooooooo.......oooooooo........oooooooooo....`,
			`.......oooooooooooooooooooooooo.......ooooooooooooooooooo.............ooooooo.........oooooooooo....`,
			`......oooooooooooooooooooooooooo.......ooooooooooooooooo...............ooooo.........oooooooooo.....`,
			`....ooooooooooooooooooooooooooooo.......ooooooooooooooooooooo......................ooooooooooo......`,
			`....ooooooooooooooooooooooooooooo.......oooooooooooooooooooooo....................oooooooooo........`,
			`......ooooooooooooooooooooooooooo.......oooooooooooooooooooooooo.................oooooooooo.........`,
			`.......oooooooooooooooooooooooooo.......oooooooooooooooooooooooooo.............ooooooooooo..........`,
			`.......ooooooooooooooooooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`......oooooooooooooooooooooooooo..........ooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`......ooooooooooooooooooooooooo..........oooooooooooooooooooooooooooooooooooooooooooooooooo.........`,
			`....oooooooooooooooooooooooooo.......:::%ooooooooooooooooooooooooooooooooooooooooooooooooo..........`,
			`....ooooooooooooooooooooooooo....:::::..ooooooooooooooooooooooooooooooooooooooooooooooooo...........`,
			`....ooooooooooooooooooooooo....:::......oooooooooooooooooooooooooooooooooooooooooooooooo............`,
			`....ooooooo...oooooooooooo%:::::.........oooooooooooooooooooooooo...............ooooooo.............`,
			`.....ooooo..........................................................................................`,
			`....................................................................................................`,
		},
	}
	cityNames = []string{
		"Abreway", "Aburg", "Accolon", "Aclomark", "Adalgar", "Adee", "Adelaide", "Adric", "Aelbarik",
		"Aeltra", "Aelwynne", "Aerilon", "Aesa", "Aethelarn", "Aeweald", "Agate City", "Agnar", "Aindorf",
		"Akyket", "Alaris", "Albrand", "Ald Ruhn", "Ald Sotha", "Aldain", "Aldcliff", "Aldmaple", "Aldmead",
		"Alera", "Alicor City", "Alldus", "Alsvider", "Alwood", "Amatin", "Anarch", "Anarth", "Andrasreth",
		"Anghurst", "Anniun", "Antbin", "Anyo", "Aodhagan", "Apore", "Appleville", "Aquarin", "Aral",
		"Aramoor", "Areth Pirn", "Argon City", "Arkaune", "Arkfire City", "Arlington", "Asbo", "Ashalmawia", "Ashland",
		"Assarinibib", "Astryde", "Atdon", "Aten", "Atlantis", "Attewelle", "Auburn", "Aumenburg", "Aure",
		"Aurion", "Axeth", "Axheim", "Ayastowe", "Ayham", "Azland", "Azmar", "Azure City", "Babrook",
		"Baelraeth", "Baeron", "Balatea", "Balderon", "Balmora", "Bamfirth", "Bampu", "Banwes", "Baraxes",
		"Barebranch", "Barghan", "Barren City", "Barrowwald", "Baseltded", "Bawside", "Bayarth", "Bayfield", "Bayrun",
		"Beacon Hill", "Befield", "Beggar’s Hole", "Belcoast", "Belgarth", "Beltran", "Belwe", "Benfield", "Benoic",
		"Benpretbrook", "Beorttun", "Berandes", "Bernthe", "Bestead", "Betport", "Bicanton", "Binacre", "Binshor",
		"Birdan", "Black Hill", "Black Hollow", "Blacksoul", "Blaglen", "Ble'eck", "Blerglust", "Blue Field", "Bluebeach",
		"Blukirk", "Blyn", "Boatwright", "Boden City", "Boheim", "Bonehorn", "Boryn", "Braddach", "Bragge",
		"Branora", "Branraker", "Branwen", "Brapslung", "Briar Glen", "Brickelwhyte", "Brine City", "Brisbane", "Bristol",
		"Briwater", "Broken Shield", "Brookmoor", "Buckmoth", "Buelthane", "Bul Isra", "Bullmar", "Buren", "Burlington",
		"Burnside", "Butterice", "Caase", "Cabury", "Caden", "Caer City", "Cafan", "Caglex", "Calanthe",
		"Caldera", "Caliron", "Calldwr", "Canberra", "Cano", "Caracatus", "Carim", "Carmisa", "Carran",
		"Carru", "Carslenford", "Cartseth", "Casthe", "Castle Hill", "Catchclaw", "Cawode", "Cear", "Cedon",
		"Cedway", "Celestra", "Cellangham", "Celoydorf", "Cengrove", "Centerville", "Cerbruk", "Cerbury", "Ceremon",
		"Cerglen", "Cerkas", "Cerrin City", "Cesirun", "Chagrad", "Chater", "Che'aldgost", "Chebas", "Chedytown",
		"Cherfren", "Chiron", "Chronos", "Cigry", "Cilburn", "Cile City", "Cinbach", "Cirrele", "Ciryon",
		"City of Fire", "City of the Shadows", "Claunecar", "Clay City", "Clayton", "Cleanbones", "Clearham Downs", "Cleveland", "Clinstan",
		"Clinton", "Clyf", "Clywd", "Coalfell", "Coel", "Coinbalth", "Coldrose", "Coldshadow", "Coldshore",
		"Condmedic", "Congisfirth", "Cormyr", "Cospera", "Cotys", "Courtmarsh", "Craeven", "Cragghe", "Crasspest",
		"Crimir", "Crystalshadow", "Cullfield", "Culshire", "Custyn", "Cyrdfel", "Cyrkarth", "Dacborath", "Dacria",
		"Daerte", "Dafarik", "Dafyd", "Dagonfel", "Dagr", "Dagrove", "Daiwick", "Dalamar", "Damarck",
		"Danfen", "Dantor", "Daon", "Darcon", "Darkbone City", "Darkmage", "Darkwell", "Darkwind", "Darwin",
		"Dasbach", "Dathoth", "Davchar", "Dayton", "Deathfall", "Dedonburn", "Dedpool", "Deepcrest", "Deephall Point",
		"Deepmoor", "Deirburn", "Delde City", "Delgrove", "Dellgate", "Delorn", "Delphys", "Denaste", "Dencede",
		"Dendest", "Dengelfel", "Deorward", "Dergost", "Derokin", "Descarq", "Detheim", "Deybank", "Dezarne",
		"Diaratyh", "Diarmaoid", "Dindale", "Dirbrand", "Disprelfield", "Distran", "Diuran", "Dogrock", "Dolgan",
		"Domin", "Doonatel", "Dorenth", "Dorhaven", "Dostborough", "Dover", "Doxca", "Doycro", "Draeden",
		"Dragonmarsh", "Drassa", "Drinishok", "Dry Gulch", "Dryope", "Duddaleah", "Dungon", "Durshire", "Dwalin",
		"Dyot", "Eadak", "Eadgyth", "Eadweard", "Earnberict", "Easthaven", "Eathelin", "Ebonheart", "Ebow City",
		"Echule", "Ecrin", "Eddra", "Edgegate", "Edorin", "Ekkel", "Eknanbor", "Elakain", "Elantir",
		"Elbramair", "Eldead", "Elderon", "Elysson", "Emerald City", "Emulpool", "Engion", "Enheim", "Enless",
		"Eoghan", "Eorforwic", "Erabenimsun", "Erast", "Erbham", "Eregdor", "Eribank", "Eron", "Esdros",
		"Esme's Rot", "Esnar", "Essault", "Estbeorn", "Estercoast", "Eststead", "Ethna", "Etnenk", "Etranth",
		"Evinob", "Faelgrar", "Faerwald", "Fairburn Point", "Fairmeadow", "Fairview", "Falconlake", "Falensarano", "Fallash Bridge",
		"Fallhedge", "Fallville", "Falo'a", "Falrepent", "Fanborough", "Fandrall", "Far Water", "Fargate", "Farnor",
		"Farshadow", "Fassen", "Faxbury", "Faycastle", "Fayfair", "Fearn", "Fearshadow", "Fedresheath", "Fegwern",
		"Feno", "Filugrave", "Firebend", "Fistrock", "Flairown", "Flat City", "Flelheim", "Fletacre", "Fondcot",
		"Fool's March", "Fornil", "Forterk", "Franklin", "Freywall", "Frostford", "Fyxinca", "Gaelen", "Gaethaa",
		"Gallys", "Gandar", "Gane", "Ganelon", "Ganith", "Garen's Well", "Garlupool", "Garn City", "Gaytforth",
		"Gealkend", "Geatan", "Geirrod", "Gentlewind", "Georgetown", "Gildeath", "Ginisis", "Gisapool", "Glasscliff",
		"Glastowe", "Glatchan", "Glingedheath", "Gnarr", "Gocin", "Godehard", "Gofannon", "Goldcrest", "Goldenleaf",
		"Goldlyn", "Goodan", "Gostarbach", "Gowerd", "Gralnen", "Gravecarn", "Greendell", "Greenhill", "Greenville",
		"Greton", "Greymage", "Greymarsh", "Greywater Edge", "Grimwall", "Grindor", "Groltain", "Gutar", "Gwayhne",
		"Haakon", "Habyn", "Hacranbrook", "Haele", "Haemfrith", "Haemin", "Haertlinde", "Hahel", "Haim",
		"Halfolk", "Halhere", "Hamish", "Haran", "Harmakros", "Harn", "Harshire", "Hasfolk", "Havale",
		"Hawkwind", "Hebost", "Hermrord", "Hestan", "Hifro", "Highdale", "Highmeadow", "Hildieth", "Hillfar",
		"Hlormaren", "Hobart", "Hogsfeet", "Holamayn", "Holith", "Hollian", "Hollyhead", "Holthasburg", "Hoochillwick",
		"Hornmar", "Hostyyk", "Hrodowald", "Hudson", "Huffimstowe", "Hull", "Hultor", "Hwen", "Hykirk",
		"Hyksos", "Iaxil", "Icebarrow", "Icefay", "Icemeet", "Iche", "Ickasu", "Illa", "Illinod",
		"Ilya", "Iniera", "Innsmouth", "Inos", "Inte", "Iprial", "Ironforge", "Ironplow", "Ironville Crossing",
		"Irragin", "Ixdencer", "Jackson", "Jancastle", "Janlyn", "Jarren's Outpost", "Jasand", "Jedarhe", "Jeling",
		"Jhena", "Jina", "Jongvale", "Jontmac", "Ka'oma", "Kara's Vale", "Kerreck", "Khartag", "Khuul",
		"Kingston", "Kior", "Kipamod", "Knife's Edge", "Koal", "Kohgoruhn", "Kouglen", "Krallides", "Kuneack",
		"Lakeshore", "Landpond", "Las Vegas", "Lassan", "Lawpest", "Laxton", "Le'oght", "Lebrus", "Leehaven",
		"Leeside", "Lenham", "Lethrys", "Lexington", "Linism", "Linland", "Linmeadow", "Lintown", "Litysh",
		"Liwald", "Lochfort", "Lochhurst", "Lookfar", "Lorbeach", "Lorbin", "Lullin", "Lyborough", "Madison",
		"Mageland", "Maire", "Mallon", "Mallowbrook", "Manchester", "Manmint", "Mantooth", "Maplehurst", "Marath",
		"Marblemoor", "Marbleton", "Margate", "Marion", "Marr Gan", "Marren's Eve", "Masellil", "Meadowlake", "Meaple",
		"Melbourne", "Merribourne", "Meshburn", "Meton", "Metropolis", "Metwan", "Mide", "Mikum", "Milford",
		"Millitburn", "Millstone", "Milton", "Mimea", "Minhanstowe", "Miranth", "Mishgrave", "Misiport", "Mitu",
		"Mompi", "Moonbright", "Moonfire", "Moonmoth", "Mora", "Morcrest", "Mount Pleasant", "Mount Vernon", "Mountmend",
		"Movawood", "Mowbach", "Muqueling", "Murplant", "Mutzcat", "Myhra", "Mysa", "Mytchville", "Nadin",
		"Nantasarn", "Narlenrun", "Nassic", "Nearon", "Nepill", "Nergwern", "Nespho", "Nestan", "Netgrove",
		"Neuson", "New Cresthill", "Newbald", "Newhaw", "Newleaf", "Newport", "Newton", "Newtown",
		"Nightfrost", "Nophalis", "Norbank", "Norbus", "Norgar", "Norratyn", "Northhollow", "Northmold", "Northpass",
		"Notlbrob", "Nuchuleft", "Nuncarth", "Nuxvar", "Nyssa", "Oakborough", "Oakheart", "Oakland", "Oaldar",
		"Oar's Rest", "Obraed", "Ociera", "Ocshire", "Odana", "Odar", "Odasgunn", "Odrosal", "Odwulf",
		"Oftar", "Old Ashton", "Oldcastle", "Oldel", "Oldshade", "Ollaneg", "Ollayos", "Olmar",
		"Olon", "Oltpest", "Omournil", "Onbruk", "Orbost", "Orfler", "Orianna", "Orness", "Orre",
		"Orrinshire", "Orwald", "Osgea", "Ossa", "Othkar", "Othon", "Oxford", "Ozorak", "Ozryn",
		"Pactra", "Pallia", "Panplara", "Pantarastar", "Pavv", "Peadar", "Peash", "Pella's Wish", "Pelra",
		"Pelthros", "Penbarn", "Penci", "Penrili", "Pentara", "Pentgaland", "Perchhead", "Perendor", "Pereswyff",
		"Perth", "Pesteir", "Petelinus", "Phames", "Piaside", "Pictar", "Pildor", "Pinnella Pass", "Pirn",
		"Plagcath", "Plemarun", "Plinsaway", "Po'asta", "Poltgobi", "Pomlinfolk", "Pontent", "Pothbaz",
		"Praice", "Pran", "Presrenfa", "Prothla", "Proudrock", "Proupol", "Prymarsh", "Prytani", "Prywyn",
		"Pultack", "Pygate", "Quagcry", "Quan Ma", "Queenstown", "Quickrock", "Raigor", "Rairkvale", "Ralsinpe",
		"Ramshorn", "Raskold", "Rathisa", "Ravenbow", "Raypond", "Rayth", "Reaver", "Red Hawk", "Redcliff",
		"Redwine", "Remdam", "Rertstet", "Resh", "Rhifirth", "Rhonius", "Riabury", "Rimmon", "Rindalsem",
		"Riveredge", "Riverside", "Riverton", "Riverwind", "Rolkfield", "Rorcy", "Roseglass", "Rosewall", "Roshun",
		"Rotshaw", "Ruanrath", "Runesward", "Rushownstad", "Rustan", "Rylla", "Ryshet", "Rysshop", "Rytor",
		"Ryvwy", "Saker Keep", "Salach", "Salamus", "Salem", "Sallen", "Saltstone", "Samca", "Sandrith",
		"Sardis", "Sarindor", "Sartheim", "Sasbank", "Scadee", "Schaldhan", "Scrushfield", "Seafort", "Seamarsh",
		"Seameet", "Secordrus", "Sedeor", "Semfirth", "Sentrun", "Serra", "Sesklos", "Sessgate", "Sestun",
		"Seyda", "Shadowdale", "Shadowmoor Downs", "Shadowpond Point", "Shavnor", "Shencodo", "Shermer", "Ship's Haven", "Shotluth",
		"Shrecrun", "Shull", "Silver City", "Silverkeep", "Silverpond Crags", "Sinyanwood", "Skyllith", "Slamer", "Sloosio",
		"Smant", "Smarmark", "Snake's Canyon", "Snowland", "Snowmelt", "Sompishaw", "Sontsil", "Sorale", "Sout",
		"South Warren", "Spedgost", "Splexnaiss", "Splustap", "Springfield", "Squall's End", "Squangrir", "Squikgost", "Stannecpan",
		"Starrycastle", "Stattown", "Steherwood", "Stenral", "Sti'yl", "Stickgate", "Stilom", "Stoneby", "Stonehand",
		"Stonle", "Stonstoke", "Storre", "Strongby", "Sudabuk", "Summerby", "Summergrass", "Summermarsh", "Sunwater",
		"Suran", "Swagrave", "Sweetwood", "Swinebroth", "Swordbreak", "Swynfield", "Sydney", "Syssale", "Tallyn",
		"Tanglen", "Taniholm", "Tarrin", "Tarvik", "Tastan", "Techanal", "Teffolk", "Telasero", "Telkna",
		"Tendaughters", "Tensa", "Tenwood", "Teran", "Tese", "Tetaun", "Thaholm", "Theeltmil", "Thorbed",
		"Thorneclay", "Thornfield", "Three Streams", "Throll", "Thyria", "Tickben", "Tijinggrad", "Tildor", "Tirne",
		"Tirport", "Tonkcent", "Toppe", "Tosdis", "Tosla", "Tramrare", "Trawald", "Trekrun", "Trion",
		"Trotwood", "Trucdon", "Trudid", "Trullion", "Tureynulal", "Twilight City", "Twisernvale", "Tywy", "Ubbin Falls",
		"Uinan", "Ula'ree", "Ulgor", "Ullast", "Unhink", "Urke", "Urshilaku", "Urvil", "Ushoul",
		"Usibel", "Ussun", "Uvirith", "Valacre", "Valenvelvar", "Valwick", "Vamstead", "Veamer", "Velishire",
		"Velothi", "Vemynal", "Veritas", "Vernolt", "Vertloch Bridge", "Vertmount Downs", "Viberg", "Vikos", "Violl's Garden",
		"Vister", "Vivec", "Vo'irnil", "Volmoria", "Vorasen", "Vril", "Wacot", "Waleoshire", "Wann",
		"Washington", "Watercoast", "Wavemeet", "Wavenhill", "Wellspring", "Wemau", "Wensol", "West Ford", "Westbay",
		"Westen", "Westertown", "Westervale", "Wetrock", "Wheatland", "Whedorf", "Whitehollow", "Whitepine", "Whiteridge",
		"Whitewell Land", "Wildebush", "Wildefort", "Wilea", "Willowdale", "Winchester", "Windrip", "Windwhisper", "Winewood",
		"Winpher", "Winterfeather", "Winterness", "Wintervale", "Wissgate", "Wistan", "Wistleigh", "Witchlyn", "Wolfden",
		"Wolfhair", "Wolfkeed", "Woodbush", "Woodend", "Woodside", "Wurcot", "Wyndu", "Xan's Bequest", "Xynnar",
		"Yarrin", "Ys", "Zaal", "Zaina", "Zao Ying", "Zapstowe", "Zeffari", "Zisbach", "Zumka",
		"Zurgonipal",
	}
)

func main() {
	for version, sourceMap := range sourceMaps {
		listOfCities := shuffleCities()

		// derive the map dimensions from the source map. maps are no longer
		// required to be a fixed size; height is the number of rows and width
		// is the length of the widest row.
		height := len(sourceMap)
		width := 0
		for _, line := range sourceMap {
			if len(line) > width {
				width = len(line)
			}
		}

		tileNames := make([][]string, height)
		cities := make([][]*City, height)
		for row := range tileNames {
			tileNames[row] = make([]string, width)
			cities[row] = make([]*City, width)
			for col := range tileNames[row] {
				tileNames[row][col] = "Blank"
			}
		}

		// map the glyphs to tile names
		for row, line := range sourceMap {
			for col, glyph := range line {
				switch glyph {
				case '#':
					tileNames[row][col] = "Blank"
				case ';':
					tileNames[row][col] = "Classic/Water Ocean Deep"
				case ',':
					tileNames[row][col] = "Classic/Water Ocean"
				case ':':
					tileNames[row][col] = "Classic/Water Sea Deep"
				case '.':
					tileNames[row][col] = "Classic/Water Sea"
				case '~':
					tileNames[row][col] = "Classic/Water Kelp"
				case ' ':
					tileNames[row][col] = "Classic/Water Kelp Heavy"
				case '"':
					tileNames[row][col] = "Classic/Water Shoals"
				case '\'':
					tileNames[row][col] = "Classic/Water Reef"
				case 'p':
					tileNames[row][col] = "Classic/Flat Grassland"
				case 'P':
					tileNames[row][col] = "Classic/Hills"
				case 'd':
					tileNames[row][col] = "Classic/Flat Desert Rocky"
				case 'D':
					tileNames[row][col] = "Classic/Flat Desert Sandy"
				case 'm':
					tileNames[row][col] = "Classic/Mountain"
				case 'M':
					tileNames[row][col] = "Classic/Mountains"
				case 's':
					tileNames[row][col] = "Classic/Flat Swamp"
				case 'S':
					tileNames[row][col] = "Classic/Flat Wetlands Jungle"
				case 'f':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
				case 'F':
					tileNames[row][col] = "Classic/Flat Forest Deciduous Heavy"
				case 'o':
					tileNames[row][col] = "Classic/Underdark Open"
				case '^':
					tileNames[row][col] = "Classic/Mountains Forest Jungle"
				case 'v':
					tileNames[row][col] = "Classic/Mountains Forest Dead"
				case '{':
					tileNames[row][col] = "Classic/Mountains Forest Deciduous"
				case '}':
					tileNames[row][col] = "Classic/Mountains Forest Evergreen"
				case ']':
					tileNames[row][col] = "Classic/Flat Marsh"
				case '[':
					tileNames[row][col] = "Classic/Flat Moor"
				case 'O':
					tileNames[row][col] = "Classic/Mountain Snowcapped"
				case '1':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Drassa", Major: true, SafeHaven: true}
				case '2':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Rimmon", Major: true, SafeHaven: true}
				case '3':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Harn", Major: true, SafeHaven: true}
				case '4':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Imperial City", Major: true, SafeHaven: true}
				case '5':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Port Aurnos", Major: true, SafeHaven: true}
				case '6':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Greyfell", Major: true, SafeHaven: true}
				case '7':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Yellowleaf", Major: true, SafeHaven: true}
				case '8':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Name: "Golden City", Major: true}
				case '9', '0':
					tileNames[row][col] = "Classic/Flat Forest Deciduous"
					cities[row][col] = &City{Major: true, SafeHaven: true}
				case '*':
					tileNames[row][col] = "Classic/Flat Grassland"
					cities[row][col] = &City{Major: true}
				case '%':
					tileNames[row][col] = "Classic/Flat Grassland"
					cities[row][col] = &City{}
				case '?':
					tileNames[row][col] = "Classic/Mountain Volcano"
				case '!':
					tileNames[row][col] = "Classic/Mountain Volcano Dormant"
				case '&':
					tileNames[row][col] = "Classic/Mountain Volcano Extinct"

				default:
					panic(fmt.Sprintf("%d:%d: %c: unknown glyph", row+1, col+1, glyph))
				}
			}
		}

		// cities
		for _, row := range cities {
			for _, city := range row {
				if city == nil {
					continue
				}
				if city.Name == "" {
					if len(listOfCities) == 0 {
						panic("assert(len(cityNames) != 0)")
					}
					city.Name = listOfCities[0]
					listOfCities = listOfCities[1:]
				}
				if city.Major {
					if city.SafeHaven {
						city.Feature = "Classic/Building Cathedral"
					} else {
						city.Feature = "Classic/Military Castle"
					}
				} else {
					if city.SafeHaven {
						city.Feature = "Classic/Building Church"
					} else {
						city.Feature = "Classic/Military Camp"
					}
				}
			}
		}

		// create a new Worldographer map with COLUMNS layout (flat-top hexes,
		// odd-q offset) and populate it from the tileNames grid.
		m := ottomap.NewMap()
		m.Name = fmt.Sprintf("Olympia G3 v%d", version+1)
		m.SetLayout(hex.OddQ)
		for row := range height {
			for col := range width {
				c := hex.FromOffset(hex.OffsetCoord{Col: col, Row: row}, hex.OddQ)
				m.SetTerrain(c, ottomap.Terrain(tileNames[row][col]))
				if city := cities[row][col]; city != nil {
					m.AddFeature(ottomap.Feature{
						Kind:     city.Feature,
						Location: c,
						Label:    city.Name,
						Layer:    ottomap.LayerAboveTerrain,
						Scale:    35,
					})
				}
			}
		}
		// pin the bounding box to the full grid.
		m.SetBounds(
			hex.FromOffset(hex.OffsetCoord{Col: 0, Row: 0}, hex.OddQ),
			hex.FromOffset(hex.OffsetCoord{Col: width - 1, Row: height - 1}, hex.OddQ),
		)

		// save the Worldographer map to cmd/g3wxx/testdata/olympia-g3-map.wxx.
		outPath := filepath.Join("cmd", "g3wxx", "testdata", fmt.Sprintf("olympia-g3-map-v%d.wxx", version+1))
		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			panic(err)
		}
		f, err := os.Create(outPath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := wog.Write(f, m, wog.WriteOptions{
			Version:     wog.V2025,
			Orientation: wog.Columns,
		}); err != nil {
			panic(err)
		}
		if err := f.Close(); err != nil {
			panic(err)
		}
		fmt.Printf("wrote %s\n", outPath)
	}
}

func shuffleCities() []string {
	list := append([]string{}, cityNames...)
	// create a random seed that isn't random for shuffling
	rnd := rand.New(rand.NewPCG(42, 42))
	rnd.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}
