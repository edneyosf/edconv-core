package aac

const codec = "libfdk_aac"
const filterChannels62 = "lowpass=c=LFE:f=120,volume=1.6,pan=stereo|FL=0.8*FL+0.5*FC+0.6*BL+0.4*LFE|FR=0.8*FR+0.5*FC+0.6*BR+0.4*LFE"

//pan=stereo| FL=0.5*FL + 0.5*FC + 0.5*SL + 0.5*LFE | FR=0.5*FR + 0.5*FC + 0.5*SR + 0.5*LFE
//pan=stereo|FL=0.8*FL+0.5*FC+0.3*SL+0.3*LFE|FR=0.8*FR+0.5*FC+0.3*SR+0.3*LFE
//lowpass=c=LFE:f=120,volume=1.6,pan=stereo|FL=0.5*FC+0.707*FL+0.707*BL+0.5*LFE|FR=0.5*FC+0.707*FR+0.707*BR+0.5*LFE