"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.ReservationInfoTypeToJSON = exports.ReservationInfoTypeFromJSONTyped = exports.ReservationInfoTypeFromJSON = exports.instanceOfReservationInfoType = void 0;
const runtime_1 = require("../runtime");
const AdvanceCheckInType_1 = require("./AdvanceCheckInType");
const CommissionPayoutToType_1 = require("./CommissionPayoutToType");
const ExternalReferenceType_1 = require("./ExternalReferenceType");
const HousekeepingRoomStatusType_1 = require("./HousekeepingRoomStatusType");
const IndicatorType_1 = require("./IndicatorType");
const PMSResStatusType_1 = require("./PMSResStatusType");
const ResAccessRestrictionType_1 = require("./ResAccessRestrictionType");
const ResAttachedProfileType_1 = require("./ResAttachedProfileType");
const ResCashieringType_1 = require("./ResCashieringType");
const ResCommunicationType_1 = require("./ResCommunicationType");
const ResGuestInfoType_1 = require("./ResGuestInfoType");
const ResHousekeepingType_1 = require("./ResHousekeepingType");
const ResMobileNotificationsType_1 = require("./ResMobileNotificationsType");
const ResRevenueBalanceType_1 = require("./ResRevenueBalanceType");
const ResSharedGuestInfoType_1 = require("./ResSharedGuestInfoType");
const ReservationAllowedActionType_1 = require("./ReservationAllowedActionType");
const ReservationDepositType_1 = require("./ReservationDepositType");
const ReservationFolioWindowType_1 = require("./ReservationFolioWindowType");
const ReservationInfoTypeCancellationInfo_1 = require("./ReservationInfoTypeCancellationInfo");
const ReservationInterfaceStatusType_1 = require("./ReservationInterfaceStatusType");
const ReservationPaymentMethodType_1 = require("./ReservationPaymentMethodType");
const ReservationQueueInformationType_1 = require("./ReservationQueueInformationType");
const ReservationTurndownInfoType_1 = require("./ReservationTurndownInfoType");
const SearchMatchType_1 = require("./SearchMatchType");
const SourceOfSaleType_1 = require("./SourceOfSaleType");
const StayInfoType_1 = require("./StayInfoType");
const TaxTypeType_1 = require("./TaxTypeType");
const UniqueIDType_1 = require("./UniqueIDType");
const WaitlistResType_1 = require("./WaitlistResType");
/**
 * Check if a given object implements the ReservationInfoType interface.
 */
function instanceOfReservationInfoType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfReservationInfoType = instanceOfReservationInfoType;
function ReservationInfoTypeFromJSON(json) {
    return ReservationInfoTypeFromJSONTyped(json, false);
}
exports.ReservationInfoTypeFromJSON = ReservationInfoTypeFromJSON;
function ReservationInfoTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'externalReferences': !(0, runtime_1.exists)(json, 'externalReferences') ? undefined : (json['externalReferences'].map(ExternalReferenceType_1.ExternalReferenceTypeFromJSON)),
        'roomStay': !(0, runtime_1.exists)(json, 'roomStay') ? undefined : (0, StayInfoType_1.StayInfoTypeFromJSON)(json['roomStay']),
        'reservationGuest': !(0, runtime_1.exists)(json, 'reservationGuest') ? undefined : (0, ResGuestInfoType_1.ResGuestInfoTypeFromJSON)(json['reservationGuest']),
        'sharedGuests': !(0, runtime_1.exists)(json, 'sharedGuests') ? undefined : (json['sharedGuests'].map(ResSharedGuestInfoType_1.ResSharedGuestInfoTypeFromJSON)),
        'attachedProfiles': !(0, runtime_1.exists)(json, 'attachedProfiles') ? undefined : (json['attachedProfiles'].map(ResAttachedProfileType_1.ResAttachedProfileTypeFromJSON)),
        'reservationPaymentMethod': !(0, runtime_1.exists)(json, 'reservationPaymentMethod') ? undefined : (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeFromJSON)(json['reservationPaymentMethod']),
        'reservationFolioWindows': !(0, runtime_1.exists)(json, 'reservationFolioWindows') ? undefined : (json['reservationFolioWindows'].map(ReservationFolioWindowType_1.ReservationFolioWindowTypeFromJSON)),
        'specials': !(0, runtime_1.exists)(json, 'specials') ? undefined : json['specials'],
        'lastPrivacyPromptDate': !(0, runtime_1.exists)(json, 'lastPrivacyPromptDate') ? undefined : (new Date(json['lastPrivacyPromptDate'])),
        'displayColor': !(0, runtime_1.exists)(json, 'displayColor') ? undefined : json['displayColor'],
        'reservationIndicators': !(0, runtime_1.exists)(json, 'reservationIndicators') ? undefined : (json['reservationIndicators'].map(IndicatorType_1.IndicatorTypeFromJSON)),
        'roomStatus': !(0, runtime_1.exists)(json, 'roomStatus') ? undefined : (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeFromJSON)(json['roomStatus']),
        'searchMatches': !(0, runtime_1.exists)(json, 'searchMatches') ? undefined : (json['searchMatches'].map(SearchMatchType_1.SearchMatchTypeFromJSON)),
        'sourceOfSale': !(0, runtime_1.exists)(json, 'sourceOfSale') ? undefined : (0, SourceOfSaleType_1.SourceOfSaleTypeFromJSON)(json['sourceOfSale']),
        'waitlist': !(0, runtime_1.exists)(json, 'waitlist') ? undefined : (0, WaitlistResType_1.WaitlistResTypeFromJSON)(json['waitlist']),
        'queue': !(0, runtime_1.exists)(json, 'queue') ? undefined : (0, ReservationQueueInformationType_1.ReservationQueueInformationTypeFromJSON)(json['queue']),
        'housekeeping': !(0, runtime_1.exists)(json, 'housekeeping') ? undefined : (0, ResHousekeepingType_1.ResHousekeepingTypeFromJSON)(json['housekeeping']),
        'cashiering': !(0, runtime_1.exists)(json, 'cashiering') ? undefined : (0, ResCashieringType_1.ResCashieringTypeFromJSON)(json['cashiering']),
        'taxType': !(0, runtime_1.exists)(json, 'taxType') ? undefined : (0, TaxTypeType_1.TaxTypeTypeFromJSON)(json['taxType']),
        'deposit': !(0, runtime_1.exists)(json, 'deposit') ? undefined : (0, ReservationDepositType_1.ReservationDepositTypeFromJSON)(json['deposit']),
        'allowedActions': !(0, runtime_1.exists)(json, 'allowedActions') ? undefined : (json['allowedActions'].map(ReservationAllowedActionType_1.ReservationAllowedActionTypeFromJSON)),
        'revenuesAndBalances': !(0, runtime_1.exists)(json, 'revenuesAndBalances') ? undefined : (0, ResRevenueBalanceType_1.ResRevenueBalanceTypeFromJSON)(json['revenuesAndBalances']),
        'hotelInterfaceStatusList': !(0, runtime_1.exists)(json, 'hotelInterfaceStatusList') ? undefined : (json['hotelInterfaceStatusList'].map(ReservationInterfaceStatusType_1.ReservationInterfaceStatusTypeFromJSON)),
        'guestPreferredCurrency': !(0, runtime_1.exists)(json, 'guestPreferredCurrency') ? undefined : json['guestPreferredCurrency'],
        'turndownInfo': !(0, runtime_1.exists)(json, 'turndownInfo') ? undefined : (0, ReservationTurndownInfoType_1.ReservationTurndownInfoTypeFromJSON)(json['turndownInfo']),
        'mobileNotifications': !(0, runtime_1.exists)(json, 'mobileNotifications') ? undefined : (0, ResMobileNotificationsType_1.ResMobileNotificationsTypeFromJSON)(json['mobileNotifications']),
        'reservationCommunication': !(0, runtime_1.exists)(json, 'reservationCommunication') ? undefined : (0, ResCommunicationType_1.ResCommunicationTypeFromJSON)(json['reservationCommunication']),
        'advanceCheckIn': !(0, runtime_1.exists)(json, 'advanceCheckIn') ? undefined : (0, AdvanceCheckInType_1.AdvanceCheckInTypeFromJSON)(json['advanceCheckIn']),
        'welcomeOffer': !(0, runtime_1.exists)(json, 'welcomeOffer') ? undefined : json['welcomeOffer'],
        'cancellationInfo': !(0, runtime_1.exists)(json, 'cancellationInfo') ? undefined : (0, ReservationInfoTypeCancellationInfo_1.ReservationInfoTypeCancellationInfoFromJSON)(json['cancellationInfo']),
        'keyCount': !(0, runtime_1.exists)(json, 'keyCount') ? undefined : json['keyCount'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'hotelName': !(0, runtime_1.exists)(json, 'hotelName') ? undefined : json['hotelName'],
        'expectedServiceTime': !(0, runtime_1.exists)(json, 'expectedServiceTime') ? undefined : json['expectedServiceTime'],
        'roomStayReservation': !(0, runtime_1.exists)(json, 'roomStayReservation') ? undefined : json['roomStayReservation'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'reservationStatus': !(0, runtime_1.exists)(json, 'reservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['reservationStatus']),
        'computedReservationStatus': !(0, runtime_1.exists)(json, 'computedReservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['computedReservationStatus']),
        'walkInIndicator': !(0, runtime_1.exists)(json, 'walkInIndicator') ? undefined : json['walkInIndicator'],
        'accessRestriction': !(0, runtime_1.exists)(json, 'accessRestriction') ? undefined : (0, ResAccessRestrictionType_1.ResAccessRestrictionTypeFromJSON)(json['accessRestriction']),
        'commissionPayoutTo': !(0, runtime_1.exists)(json, 'commissionPayoutTo') ? undefined : (0, CommissionPayoutToType_1.CommissionPayoutToTypeFromJSON)(json['commissionPayoutTo']),
        'paymentMethod': !(0, runtime_1.exists)(json, 'paymentMethod') ? undefined : json['paymentMethod'],
        'preRegistered': !(0, runtime_1.exists)(json, 'preRegistered') ? undefined : json['preRegistered'],
        'openFolio': !(0, runtime_1.exists)(json, 'openFolio') ? undefined : json['openFolio'],
        'allowMobileCheckout': !(0, runtime_1.exists)(json, 'allowMobileCheckout') ? undefined : json['allowMobileCheckout'],
        'allowMobileViewFolio': !(0, runtime_1.exists)(json, 'allowMobileViewFolio') ? undefined : json['allowMobileViewFolio'],
        'optedForCommunication': !(0, runtime_1.exists)(json, 'optedForCommunication') ? undefined : json['optedForCommunication'],
    };
}
exports.ReservationInfoTypeFromJSONTyped = ReservationInfoTypeFromJSONTyped;
function ReservationInfoTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'externalReferences': value.externalReferences === undefined ? undefined : (value.externalReferences.map(ExternalReferenceType_1.ExternalReferenceTypeToJSON)),
        'roomStay': (0, StayInfoType_1.StayInfoTypeToJSON)(value.roomStay),
        'reservationGuest': (0, ResGuestInfoType_1.ResGuestInfoTypeToJSON)(value.reservationGuest),
        'sharedGuests': value.sharedGuests === undefined ? undefined : (value.sharedGuests.map(ResSharedGuestInfoType_1.ResSharedGuestInfoTypeToJSON)),
        'attachedProfiles': value.attachedProfiles === undefined ? undefined : (value.attachedProfiles.map(ResAttachedProfileType_1.ResAttachedProfileTypeToJSON)),
        'reservationPaymentMethod': (0, ReservationPaymentMethodType_1.ReservationPaymentMethodTypeToJSON)(value.reservationPaymentMethod),
        'reservationFolioWindows': value.reservationFolioWindows === undefined ? undefined : (value.reservationFolioWindows.map(ReservationFolioWindowType_1.ReservationFolioWindowTypeToJSON)),
        'specials': value.specials,
        'lastPrivacyPromptDate': value.lastPrivacyPromptDate === undefined ? undefined : (value.lastPrivacyPromptDate.toISOString().substr(0, 10)),
        'displayColor': value.displayColor,
        'reservationIndicators': value.reservationIndicators === undefined ? undefined : (value.reservationIndicators.map(IndicatorType_1.IndicatorTypeToJSON)),
        'roomStatus': (0, HousekeepingRoomStatusType_1.HousekeepingRoomStatusTypeToJSON)(value.roomStatus),
        'searchMatches': value.searchMatches === undefined ? undefined : (value.searchMatches.map(SearchMatchType_1.SearchMatchTypeToJSON)),
        'sourceOfSale': (0, SourceOfSaleType_1.SourceOfSaleTypeToJSON)(value.sourceOfSale),
        'waitlist': (0, WaitlistResType_1.WaitlistResTypeToJSON)(value.waitlist),
        'queue': (0, ReservationQueueInformationType_1.ReservationQueueInformationTypeToJSON)(value.queue),
        'housekeeping': (0, ResHousekeepingType_1.ResHousekeepingTypeToJSON)(value.housekeeping),
        'cashiering': (0, ResCashieringType_1.ResCashieringTypeToJSON)(value.cashiering),
        'taxType': (0, TaxTypeType_1.TaxTypeTypeToJSON)(value.taxType),
        'deposit': (0, ReservationDepositType_1.ReservationDepositTypeToJSON)(value.deposit),
        'allowedActions': value.allowedActions === undefined ? undefined : (value.allowedActions.map(ReservationAllowedActionType_1.ReservationAllowedActionTypeToJSON)),
        'revenuesAndBalances': (0, ResRevenueBalanceType_1.ResRevenueBalanceTypeToJSON)(value.revenuesAndBalances),
        'hotelInterfaceStatusList': value.hotelInterfaceStatusList === undefined ? undefined : (value.hotelInterfaceStatusList.map(ReservationInterfaceStatusType_1.ReservationInterfaceStatusTypeToJSON)),
        'guestPreferredCurrency': value.guestPreferredCurrency,
        'turndownInfo': (0, ReservationTurndownInfoType_1.ReservationTurndownInfoTypeToJSON)(value.turndownInfo),
        'mobileNotifications': (0, ResMobileNotificationsType_1.ResMobileNotificationsTypeToJSON)(value.mobileNotifications),
        'reservationCommunication': (0, ResCommunicationType_1.ResCommunicationTypeToJSON)(value.reservationCommunication),
        'advanceCheckIn': (0, AdvanceCheckInType_1.AdvanceCheckInTypeToJSON)(value.advanceCheckIn),
        'welcomeOffer': value.welcomeOffer,
        'cancellationInfo': (0, ReservationInfoTypeCancellationInfo_1.ReservationInfoTypeCancellationInfoToJSON)(value.cancellationInfo),
        'keyCount': value.keyCount,
        'hotelId': value.hotelId,
        'hotelName': value.hotelName,
        'expectedServiceTime': value.expectedServiceTime,
        'roomStayReservation': value.roomStayReservation,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'reservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.reservationStatus),
        'computedReservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.computedReservationStatus),
        'walkInIndicator': value.walkInIndicator,
        'accessRestriction': (0, ResAccessRestrictionType_1.ResAccessRestrictionTypeToJSON)(value.accessRestriction),
        'commissionPayoutTo': (0, CommissionPayoutToType_1.CommissionPayoutToTypeToJSON)(value.commissionPayoutTo),
        'paymentMethod': value.paymentMethod,
        'preRegistered': value.preRegistered,
        'openFolio': value.openFolio,
        'allowMobileCheckout': value.allowMobileCheckout,
        'allowMobileViewFolio': value.allowMobileViewFolio,
        'optedForCommunication': value.optedForCommunication,
    };
}
exports.ReservationInfoTypeToJSON = ReservationInfoTypeToJSON;
