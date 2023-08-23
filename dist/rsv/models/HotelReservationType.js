"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.HotelReservationTypeToJSON = exports.HotelReservationTypeFromJSONTyped = exports.HotelReservationTypeFromJSON = exports.instanceOfHotelReservationType = void 0;
const runtime_1 = require("../runtime");
const AdvanceCheckInType_1 = require("./AdvanceCheckInType");
const AlertType_1 = require("./AlertType");
const AttachmentType_1 = require("./AttachmentType");
const AwardType_1 = require("./AwardType");
const CallType_1 = require("./CallType");
const CateringResInfoType_1 = require("./CateringResInfoType");
const CommentInfoType_1 = require("./CommentInfoType");
const CompAuthorizerInfoType_1 = require("./CompAuthorizerInfoType");
const ConfirmationType_1 = require("./ConfirmationType");
const CustomChargeExemptionType_1 = require("./CustomChargeExemptionType");
const ECouponType_1 = require("./ECouponType");
const ExternalReferenceType_1 = require("./ExternalReferenceType");
const FetchActivityBookingsType_1 = require("./FetchActivityBookingsType");
const FixedChargeType_1 = require("./FixedChargeType");
const GuestMessageType_1 = require("./GuestMessageType");
const HotelReservationTraceType_1 = require("./HotelReservationTraceType");
const HotelReservationTypeCancellation_1 = require("./HotelReservationTypeCancellation");
const HotelReservationTypeReservationProfiles_1 = require("./HotelReservationTypeReservationProfiles");
const HotelReservationTypeTransactionDiversions_1 = require("./HotelReservationTypeTransactionDiversions");
const HotelReservationsType_1 = require("./HotelReservationsType");
const IndicatorType_1 = require("./IndicatorType");
const LinkedReservationsInfoType_1 = require("./LinkedReservationsInfoType");
const MembershipType_1 = require("./MembershipType");
const NameValueDetailType_1 = require("./NameValueDetailType");
const OverrideInstructionType_1 = require("./OverrideInstructionType");
const PMSResStatusType_1 = require("./PMSResStatusType");
const PreferenceTypeType_1 = require("./PreferenceTypeType");
const PrepaidCardType_1 = require("./PrepaidCardType");
const ResAccessRestrictionType_1 = require("./ResAccessRestrictionType");
const ResCashieringType_1 = require("./ResCashieringType");
const ResCommunicationType_1 = require("./ResCommunicationType");
const ResGuestAdditionalInfoType_1 = require("./ResGuestAdditionalInfoType");
const ResGuestType_1 = require("./ResGuestType");
const ResHousekeepingType_1 = require("./ResHousekeepingType");
const ResInventoryItemsType_1 = require("./ResInventoryItemsType");
const ResSharedGuestInfoType_1 = require("./ResSharedGuestInfoType");
const ReservationAllowedActionType_1 = require("./ReservationAllowedActionType");
const ReservationECertificateType_1 = require("./ReservationECertificateType");
const ReservationLocatorType_1 = require("./ReservationLocatorType");
const ReservationMembershipType_1 = require("./ReservationMembershipType");
const ReservationPackageType_1 = require("./ReservationPackageType");
const ReservationPaymentMethodType_1 = require("./ReservationPaymentMethodType");
const ReservationPoliciesType_1 = require("./ReservationPoliciesType");
const ReservationQueueInformationType_1 = require("./ReservationQueueInformationType");
const RoomStayType_1 = require("./RoomStayType");
const RoutingInfoType_1 = require("./RoutingInfoType");
const ServiceRequest_1 = require("./ServiceRequest");
const SourceOfSaleType_1 = require("./SourceOfSaleType");
const TicketType_1 = require("./TicketType");
const TrackItItemType_1 = require("./TrackItItemType");
const UniqueIDType_1 = require("./UniqueIDType");
const UserDefinedFieldsType_1 = require("./UserDefinedFieldsType");
const WaitlistResType_1 = require("./WaitlistResType");
/**
 * Check if a given object implements the HotelReservationType interface.
 */
function instanceOfHotelReservationType(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfHotelReservationType = instanceOfHotelReservationType;
function HotelReservationTypeFromJSON(json) {
    return HotelReservationTypeFromJSONTyped(json, false);
}
exports.HotelReservationTypeFromJSON = HotelReservationTypeFromJSON;
function HotelReservationTypeFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'reservationIdList': !(0, runtime_1.exists)(json, 'reservationIdList') ? undefined : (json['reservationIdList'].map(UniqueIDType_1.UniqueIDTypeFromJSON)),
        'externalReferences': !(0, runtime_1.exists)(json, 'externalReferences') ? undefined : (json['externalReferences'].map(ExternalReferenceType_1.ExternalReferenceTypeFromJSON)),
        'sourceOfSale': !(0, runtime_1.exists)(json, 'sourceOfSale') ? undefined : (0, SourceOfSaleType_1.SourceOfSaleTypeFromJSON)(json['sourceOfSale']),
        'roomStay': !(0, runtime_1.exists)(json, 'roomStay') ? undefined : (0, RoomStayType_1.RoomStayTypeFromJSON)(json['roomStay']),
        'compAuthorizer': !(0, runtime_1.exists)(json, 'compAuthorizer') ? undefined : (0, CompAuthorizerInfoType_1.CompAuthorizerInfoTypeFromJSON)(json['compAuthorizer']),
        'reservationGuests': !(0, runtime_1.exists)(json, 'reservationGuests') ? undefined : (json['reservationGuests'].map(ResGuestType_1.ResGuestTypeFromJSON)),
        'sharedGuests': !(0, runtime_1.exists)(json, 'sharedGuests') ? undefined : (json['sharedGuests'].map(ResSharedGuestInfoType_1.ResSharedGuestInfoTypeFromJSON)),
        'additionalGuestInfo': !(0, runtime_1.exists)(json, 'additionalGuestInfo') ? undefined : (0, ResGuestAdditionalInfoType_1.ResGuestAdditionalInfoTypeFromJSON)(json['additionalGuestInfo']),
        'reservationProfiles': !(0, runtime_1.exists)(json, 'reservationProfiles') ? undefined : (0, HotelReservationTypeReservationProfiles_1.HotelReservationTypeReservationProfilesFromJSON)(json['reservationProfiles']),
        'reservationCommunication': !(0, runtime_1.exists)(json, 'reservationCommunication') ? undefined : (0, ResCommunicationType_1.ResCommunicationTypeFromJSON)(json['reservationCommunication']),
        'trackItItems': !(0, runtime_1.exists)(json, 'trackItItems') ? undefined : (json['trackItItems'].map(TrackItItemType_1.TrackItItemTypeFromJSON)),
        'shares': !(0, runtime_1.exists)(json, 'shares') ? undefined : (0, HotelReservationsType_1.HotelReservationsTypeFromJSON)(json['shares']),
        'locators': !(0, runtime_1.exists)(json, 'locators') ? undefined : (json['locators'].map(ReservationLocatorType_1.ReservationLocatorTypeFromJSON)),
        'attachments': !(0, runtime_1.exists)(json, 'attachments') ? undefined : (json['attachments'].map(AttachmentType_1.AttachmentTypeFromJSON)),
        'webRegistrationCards': !(0, runtime_1.exists)(json, 'webRegistrationCards') ? undefined : (json['webRegistrationCards'].map(AttachmentType_1.AttachmentTypeFromJSON)),
        'serviceRequests': !(0, runtime_1.exists)(json, 'serviceRequests') ? undefined : (json['serviceRequests'].map(ServiceRequest_1.ServiceRequestFromJSON)),
        'reservationActivities': !(0, runtime_1.exists)(json, 'reservationActivities') ? undefined : (json['reservationActivities'].map(FetchActivityBookingsType_1.FetchActivityBookingsTypeFromJSON)),
        'scheduledActivities': !(0, runtime_1.exists)(json, 'scheduledActivities') ? undefined : (json['scheduledActivities'].map(FetchActivityBookingsType_1.FetchActivityBookingsTypeFromJSON)),
        'prepaidCards': !(0, runtime_1.exists)(json, 'prepaidCards') ? undefined : (json['prepaidCards'].map(PrepaidCardType_1.PrepaidCardTypeFromJSON)),
        'profileAwards': !(0, runtime_1.exists)(json, 'profileAwards') ? undefined : (json['profileAwards'].map(AwardType_1.AwardTypeFromJSON)),
        'reservationPackages': !(0, runtime_1.exists)(json, 'reservationPackages') ? undefined : (json['reservationPackages'].map(ReservationPackageType_1.ReservationPackageTypeFromJSON)),
        'inventoryItems': !(0, runtime_1.exists)(json, 'inventoryItems') ? undefined : (0, ResInventoryItemsType_1.ResInventoryItemsTypeFromJSON)(json['inventoryItems']),
        'comments': !(0, runtime_1.exists)(json, 'comments') ? undefined : (json['comments'].map(CommentInfoType_1.CommentInfoTypeFromJSON)),
        'guestComments': !(0, runtime_1.exists)(json, 'guestComments') ? undefined : (json['guestComments'].map(CommentInfoType_1.CommentInfoTypeFromJSON)),
        'guestMemberships': !(0, runtime_1.exists)(json, 'guestMemberships') ? undefined : (json['guestMemberships'].map(MembershipType_1.MembershipTypeFromJSON)),
        'preferenceCollection': !(0, runtime_1.exists)(json, 'preferenceCollection') ? undefined : (json['preferenceCollection'].map(PreferenceTypeType_1.PreferenceTypeTypeFromJSON)),
        'reservationMemberships': !(0, runtime_1.exists)(json, 'reservationMemberships') ? undefined : (json['reservationMemberships'].map(ReservationMembershipType_1.ReservationMembershipTypeFromJSON)),
        'reservationPaymentMethods': !(0, runtime_1.exists)(json, 'reservationPaymentMethods') ? undefined : (json['reservationPaymentMethods'].map(ReservationPaymentMethodType_1.ReservationPaymentMethodTypeFromJSON)),
        'routingInstructions': !(0, runtime_1.exists)(json, 'routingInstructions') ? undefined : (json['routingInstructions'].map(RoutingInfoType_1.RoutingInfoTypeFromJSON)),
        'reservationPolicies': !(0, runtime_1.exists)(json, 'reservationPolicies') ? undefined : (0, ReservationPoliciesType_1.ReservationPoliciesTypeFromJSON)(json['reservationPolicies']),
        'cashiering': !(0, runtime_1.exists)(json, 'cashiering') ? undefined : (0, ResCashieringType_1.ResCashieringTypeFromJSON)(json['cashiering']),
        'housekeeping': !(0, runtime_1.exists)(json, 'housekeeping') ? undefined : (0, ResHousekeepingType_1.ResHousekeepingTypeFromJSON)(json['housekeeping']),
        'linkedReservation': !(0, runtime_1.exists)(json, 'linkedReservation') ? undefined : (0, LinkedReservationsInfoType_1.LinkedReservationsInfoTypeFromJSON)(json['linkedReservation']),
        'extSystemSync': !(0, runtime_1.exists)(json, 'extSystemSync') ? undefined : json['extSystemSync'],
        'userDefinedFields': !(0, runtime_1.exists)(json, 'userDefinedFields') ? undefined : (0, UserDefinedFieldsType_1.UserDefinedFieldsTypeFromJSON)(json['userDefinedFields']),
        'reservationIndicators': !(0, runtime_1.exists)(json, 'reservationIndicators') ? undefined : (json['reservationIndicators'].map(IndicatorType_1.IndicatorTypeFromJSON)),
        'waitlist': !(0, runtime_1.exists)(json, 'waitlist') ? undefined : (0, WaitlistResType_1.WaitlistResTypeFromJSON)(json['waitlist']),
        'cancellation': !(0, runtime_1.exists)(json, 'cancellation') ? undefined : (0, HotelReservationTypeCancellation_1.HotelReservationTypeCancellationFromJSON)(json['cancellation']),
        'catering': !(0, runtime_1.exists)(json, 'catering') ? undefined : (0, CateringResInfoType_1.CateringResInfoTypeFromJSON)(json['catering']),
        'alerts': !(0, runtime_1.exists)(json, 'alerts') ? undefined : (json['alerts'].map(AlertType_1.AlertTypeFromJSON)),
        'traces': !(0, runtime_1.exists)(json, 'traces') ? undefined : (json['traces'].map(HotelReservationTraceType_1.HotelReservationTraceTypeFromJSON)),
        'confirmationLetters': !(0, runtime_1.exists)(json, 'confirmationLetters') ? undefined : (json['confirmationLetters'].map(ConfirmationType_1.ConfirmationTypeFromJSON)),
        'callHistory': !(0, runtime_1.exists)(json, 'callHistory') ? undefined : (json['callHistory'].map(CallType_1.CallTypeFromJSON)),
        'fixedCharges': !(0, runtime_1.exists)(json, 'fixedCharges') ? undefined : (json['fixedCharges'].map(FixedChargeType_1.FixedChargeTypeFromJSON)),
        'guestMessages': !(0, runtime_1.exists)(json, 'guestMessages') ? undefined : (json['guestMessages'].map(GuestMessageType_1.GuestMessageTypeFromJSON)),
        'lockHandle': !(0, runtime_1.exists)(json, 'lockHandle') ? undefined : json['lockHandle'],
        'overrideInstructions': !(0, runtime_1.exists)(json, 'overrideInstructions') ? undefined : (json['overrideInstructions'].map(OverrideInstructionType_1.OverrideInstructionTypeFromJSON)),
        'queue': !(0, runtime_1.exists)(json, 'queue') ? undefined : (0, ReservationQueueInformationType_1.ReservationQueueInformationTypeFromJSON)(json['queue']),
        'allowedActions': !(0, runtime_1.exists)(json, 'allowedActions') ? undefined : (json['allowedActions'].map(ReservationAllowedActionType_1.ReservationAllowedActionTypeFromJSON)),
        'eCoupons': !(0, runtime_1.exists)(json, 'eCoupons') ? undefined : (json['eCoupons'].map(ECouponType_1.ECouponTypeFromJSON)),
        'transactionDiversions': !(0, runtime_1.exists)(json, 'transactionDiversions') ? undefined : (0, HotelReservationTypeTransactionDiversions_1.HotelReservationTypeTransactionDiversionsFromJSON)(json['transactionDiversions']),
        'advanceCheckIn': !(0, runtime_1.exists)(json, 'advanceCheckIn') ? undefined : (0, AdvanceCheckInType_1.AdvanceCheckInTypeFromJSON)(json['advanceCheckIn']),
        'tickets': !(0, runtime_1.exists)(json, 'tickets') ? undefined : (json['tickets'].map(TicketType_1.TicketTypeFromJSON)),
        'accessRestrictionMessage': !(0, runtime_1.exists)(json, 'accessRestrictionMessage') ? undefined : json['accessRestrictionMessage'],
        'eCertificates': !(0, runtime_1.exists)(json, 'eCertificates') ? undefined : (json['eCertificates'].map(ReservationECertificateType_1.ReservationECertificateTypeFromJSON)),
        'customNameValueDetail': !(0, runtime_1.exists)(json, 'customNameValueDetail') ? undefined : (0, NameValueDetailType_1.NameValueDetailTypeFromJSON)(json['customNameValueDetail']),
        'customChargeExemptionDetails': !(0, runtime_1.exists)(json, 'customChargeExemptionDetails') ? undefined : (json['customChargeExemptionDetails'].map(CustomChargeExemptionType_1.CustomChargeExemptionTypeFromJSON)),
        'autoBorrowFromHouse': !(0, runtime_1.exists)(json, 'autoBorrowFromHouse') ? undefined : json['autoBorrowFromHouse'],
        'overrideExternalChecks': !(0, runtime_1.exists)(json, 'overrideExternalChecks') ? undefined : json['overrideExternalChecks'],
        'hotelId': !(0, runtime_1.exists)(json, 'hotelId') ? undefined : json['hotelId'],
        'roomStayReservation': !(0, runtime_1.exists)(json, 'roomStayReservation') ? undefined : json['roomStayReservation'],
        'reservationStatus': !(0, runtime_1.exists)(json, 'reservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['reservationStatus']),
        'computedReservationStatus': !(0, runtime_1.exists)(json, 'computedReservationStatus') ? undefined : (0, PMSResStatusType_1.PMSResStatusTypeFromJSON)(json['computedReservationStatus']),
        'walkIn': !(0, runtime_1.exists)(json, 'walkIn') ? undefined : json['walkIn'],
        'printRate': !(0, runtime_1.exists)(json, 'printRate') ? undefined : json['printRate'],
        'createDateTime': !(0, runtime_1.exists)(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !(0, runtime_1.exists)(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !(0, runtime_1.exists)(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !(0, runtime_1.exists)(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'purgeDate': !(0, runtime_1.exists)(json, 'purgeDate') ? undefined : (new Date(json['purgeDate'])),
        'createBusinessDate': !(0, runtime_1.exists)(json, 'createBusinessDate') ? undefined : (new Date(json['createBusinessDate'])),
        'reinstateDate': !(0, runtime_1.exists)(json, 'reinstateDate') ? undefined : (new Date(json['reinstateDate'])),
        'party': !(0, runtime_1.exists)(json, 'party') ? undefined : json['party'],
        'primaryEnrollmentResort': !(0, runtime_1.exists)(json, 'primaryEnrollmentResort') ? undefined : json['primaryEnrollmentResort'],
        'primaryEnrollmentChain': !(0, runtime_1.exists)(json, 'primaryEnrollmentChain') ? undefined : json['primaryEnrollmentChain'],
        'customReference': !(0, runtime_1.exists)(json, 'customReference') ? undefined : json['customReference'],
        'displayColor': !(0, runtime_1.exists)(json, 'displayColor') ? undefined : json['displayColor'],
        'markAsRecentlyAccessed': !(0, runtime_1.exists)(json, 'markAsRecentlyAccessed') ? undefined : json['markAsRecentlyAccessed'],
        'overrideInventoryCheck': !(0, runtime_1.exists)(json, 'overrideInventoryCheck') ? undefined : json['overrideInventoryCheck'],
        'accessRestriction': !(0, runtime_1.exists)(json, 'accessRestriction') ? undefined : (0, ResAccessRestrictionType_1.ResAccessRestrictionTypeFromJSON)(json['accessRestriction']),
        'preRegistered': !(0, runtime_1.exists)(json, 'preRegistered') ? undefined : json['preRegistered'],
        'upgradeEligible': !(0, runtime_1.exists)(json, 'upgradeEligible') ? undefined : json['upgradeEligible'],
        'overrideBlockRestriction': !(0, runtime_1.exists)(json, 'overrideBlockRestriction') ? undefined : json['overrideBlockRestriction'],
        'allowAutoCheckin': !(0, runtime_1.exists)(json, 'allowAutoCheckin') ? undefined : json['allowAutoCheckin'],
        'hasOpenFolio': !(0, runtime_1.exists)(json, 'hasOpenFolio') ? undefined : json['hasOpenFolio'],
        'allowMobileCheckout': !(0, runtime_1.exists)(json, 'allowMobileCheckout') ? undefined : json['allowMobileCheckout'],
        'allowMobileViewFolio': !(0, runtime_1.exists)(json, 'allowMobileViewFolio') ? undefined : json['allowMobileViewFolio'],
        'allowPreRegistration': !(0, runtime_1.exists)(json, 'allowPreRegistration') ? undefined : json['allowPreRegistration'],
        'optedForCommunication': !(0, runtime_1.exists)(json, 'optedForCommunication') ? undefined : json['optedForCommunication'],
        'chargeCardNumber': !(0, runtime_1.exists)(json, 'chargeCardNumber') ? undefined : json['chargeCardNumber'],
    };
}
exports.HotelReservationTypeFromJSONTyped = HotelReservationTypeFromJSONTyped;
function HotelReservationTypeToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'reservationIdList': value.reservationIdList === undefined ? undefined : (value.reservationIdList.map(UniqueIDType_1.UniqueIDTypeToJSON)),
        'externalReferences': value.externalReferences === undefined ? undefined : (value.externalReferences.map(ExternalReferenceType_1.ExternalReferenceTypeToJSON)),
        'sourceOfSale': (0, SourceOfSaleType_1.SourceOfSaleTypeToJSON)(value.sourceOfSale),
        'roomStay': (0, RoomStayType_1.RoomStayTypeToJSON)(value.roomStay),
        'compAuthorizer': (0, CompAuthorizerInfoType_1.CompAuthorizerInfoTypeToJSON)(value.compAuthorizer),
        'reservationGuests': value.reservationGuests === undefined ? undefined : (value.reservationGuests.map(ResGuestType_1.ResGuestTypeToJSON)),
        'sharedGuests': value.sharedGuests === undefined ? undefined : (value.sharedGuests.map(ResSharedGuestInfoType_1.ResSharedGuestInfoTypeToJSON)),
        'additionalGuestInfo': (0, ResGuestAdditionalInfoType_1.ResGuestAdditionalInfoTypeToJSON)(value.additionalGuestInfo),
        'reservationProfiles': (0, HotelReservationTypeReservationProfiles_1.HotelReservationTypeReservationProfilesToJSON)(value.reservationProfiles),
        'reservationCommunication': (0, ResCommunicationType_1.ResCommunicationTypeToJSON)(value.reservationCommunication),
        'trackItItems': value.trackItItems === undefined ? undefined : (value.trackItItems.map(TrackItItemType_1.TrackItItemTypeToJSON)),
        'shares': (0, HotelReservationsType_1.HotelReservationsTypeToJSON)(value.shares),
        'locators': value.locators === undefined ? undefined : (value.locators.map(ReservationLocatorType_1.ReservationLocatorTypeToJSON)),
        'attachments': value.attachments === undefined ? undefined : (value.attachments.map(AttachmentType_1.AttachmentTypeToJSON)),
        'webRegistrationCards': value.webRegistrationCards === undefined ? undefined : (value.webRegistrationCards.map(AttachmentType_1.AttachmentTypeToJSON)),
        'serviceRequests': value.serviceRequests === undefined ? undefined : (value.serviceRequests.map(ServiceRequest_1.ServiceRequestToJSON)),
        'reservationActivities': value.reservationActivities === undefined ? undefined : (value.reservationActivities.map(FetchActivityBookingsType_1.FetchActivityBookingsTypeToJSON)),
        'scheduledActivities': value.scheduledActivities === undefined ? undefined : (value.scheduledActivities.map(FetchActivityBookingsType_1.FetchActivityBookingsTypeToJSON)),
        'prepaidCards': value.prepaidCards === undefined ? undefined : (value.prepaidCards.map(PrepaidCardType_1.PrepaidCardTypeToJSON)),
        'profileAwards': value.profileAwards === undefined ? undefined : (value.profileAwards.map(AwardType_1.AwardTypeToJSON)),
        'reservationPackages': value.reservationPackages === undefined ? undefined : (value.reservationPackages.map(ReservationPackageType_1.ReservationPackageTypeToJSON)),
        'inventoryItems': (0, ResInventoryItemsType_1.ResInventoryItemsTypeToJSON)(value.inventoryItems),
        'comments': value.comments === undefined ? undefined : (value.comments.map(CommentInfoType_1.CommentInfoTypeToJSON)),
        'guestComments': value.guestComments === undefined ? undefined : (value.guestComments.map(CommentInfoType_1.CommentInfoTypeToJSON)),
        'guestMemberships': value.guestMemberships === undefined ? undefined : (value.guestMemberships.map(MembershipType_1.MembershipTypeToJSON)),
        'preferenceCollection': value.preferenceCollection === undefined ? undefined : (value.preferenceCollection.map(PreferenceTypeType_1.PreferenceTypeTypeToJSON)),
        'reservationMemberships': value.reservationMemberships === undefined ? undefined : (value.reservationMemberships.map(ReservationMembershipType_1.ReservationMembershipTypeToJSON)),
        'reservationPaymentMethods': value.reservationPaymentMethods === undefined ? undefined : (value.reservationPaymentMethods.map(ReservationPaymentMethodType_1.ReservationPaymentMethodTypeToJSON)),
        'routingInstructions': value.routingInstructions === undefined ? undefined : (value.routingInstructions.map(RoutingInfoType_1.RoutingInfoTypeToJSON)),
        'reservationPolicies': (0, ReservationPoliciesType_1.ReservationPoliciesTypeToJSON)(value.reservationPolicies),
        'cashiering': (0, ResCashieringType_1.ResCashieringTypeToJSON)(value.cashiering),
        'housekeeping': (0, ResHousekeepingType_1.ResHousekeepingTypeToJSON)(value.housekeeping),
        'linkedReservation': (0, LinkedReservationsInfoType_1.LinkedReservationsInfoTypeToJSON)(value.linkedReservation),
        'extSystemSync': value.extSystemSync,
        'userDefinedFields': (0, UserDefinedFieldsType_1.UserDefinedFieldsTypeToJSON)(value.userDefinedFields),
        'reservationIndicators': value.reservationIndicators === undefined ? undefined : (value.reservationIndicators.map(IndicatorType_1.IndicatorTypeToJSON)),
        'waitlist': (0, WaitlistResType_1.WaitlistResTypeToJSON)(value.waitlist),
        'cancellation': (0, HotelReservationTypeCancellation_1.HotelReservationTypeCancellationToJSON)(value.cancellation),
        'catering': (0, CateringResInfoType_1.CateringResInfoTypeToJSON)(value.catering),
        'alerts': value.alerts === undefined ? undefined : (value.alerts.map(AlertType_1.AlertTypeToJSON)),
        'traces': value.traces === undefined ? undefined : (value.traces.map(HotelReservationTraceType_1.HotelReservationTraceTypeToJSON)),
        'confirmationLetters': value.confirmationLetters === undefined ? undefined : (value.confirmationLetters.map(ConfirmationType_1.ConfirmationTypeToJSON)),
        'callHistory': value.callHistory === undefined ? undefined : (value.callHistory.map(CallType_1.CallTypeToJSON)),
        'fixedCharges': value.fixedCharges === undefined ? undefined : (value.fixedCharges.map(FixedChargeType_1.FixedChargeTypeToJSON)),
        'guestMessages': value.guestMessages === undefined ? undefined : (value.guestMessages.map(GuestMessageType_1.GuestMessageTypeToJSON)),
        'lockHandle': value.lockHandle,
        'overrideInstructions': value.overrideInstructions === undefined ? undefined : (value.overrideInstructions.map(OverrideInstructionType_1.OverrideInstructionTypeToJSON)),
        'queue': (0, ReservationQueueInformationType_1.ReservationQueueInformationTypeToJSON)(value.queue),
        'allowedActions': value.allowedActions === undefined ? undefined : (value.allowedActions.map(ReservationAllowedActionType_1.ReservationAllowedActionTypeToJSON)),
        'eCoupons': value.eCoupons === undefined ? undefined : (value.eCoupons.map(ECouponType_1.ECouponTypeToJSON)),
        'transactionDiversions': (0, HotelReservationTypeTransactionDiversions_1.HotelReservationTypeTransactionDiversionsToJSON)(value.transactionDiversions),
        'advanceCheckIn': (0, AdvanceCheckInType_1.AdvanceCheckInTypeToJSON)(value.advanceCheckIn),
        'tickets': value.tickets === undefined ? undefined : (value.tickets.map(TicketType_1.TicketTypeToJSON)),
        'accessRestrictionMessage': value.accessRestrictionMessage,
        'eCertificates': value.eCertificates === undefined ? undefined : (value.eCertificates.map(ReservationECertificateType_1.ReservationECertificateTypeToJSON)),
        'customNameValueDetail': (0, NameValueDetailType_1.NameValueDetailTypeToJSON)(value.customNameValueDetail),
        'customChargeExemptionDetails': value.customChargeExemptionDetails === undefined ? undefined : (value.customChargeExemptionDetails.map(CustomChargeExemptionType_1.CustomChargeExemptionTypeToJSON)),
        'autoBorrowFromHouse': value.autoBorrowFromHouse,
        'overrideExternalChecks': value.overrideExternalChecks,
        'hotelId': value.hotelId,
        'roomStayReservation': value.roomStayReservation,
        'reservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.reservationStatus),
        'computedReservationStatus': (0, PMSResStatusType_1.PMSResStatusTypeToJSON)(value.computedReservationStatus),
        'walkIn': value.walkIn,
        'printRate': value.printRate,
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'purgeDate': value.purgeDate === undefined ? undefined : (value.purgeDate.toISOString().substr(0, 10)),
        'createBusinessDate': value.createBusinessDate === undefined ? undefined : (value.createBusinessDate.toISOString().substr(0, 10)),
        'reinstateDate': value.reinstateDate === undefined ? undefined : (value.reinstateDate.toISOString().substr(0, 10)),
        'party': value.party,
        'primaryEnrollmentResort': value.primaryEnrollmentResort,
        'primaryEnrollmentChain': value.primaryEnrollmentChain,
        'customReference': value.customReference,
        'displayColor': value.displayColor,
        'markAsRecentlyAccessed': value.markAsRecentlyAccessed,
        'overrideInventoryCheck': value.overrideInventoryCheck,
        'accessRestriction': (0, ResAccessRestrictionType_1.ResAccessRestrictionTypeToJSON)(value.accessRestriction),
        'preRegistered': value.preRegistered,
        'upgradeEligible': value.upgradeEligible,
        'overrideBlockRestriction': value.overrideBlockRestriction,
        'allowAutoCheckin': value.allowAutoCheckin,
        'hasOpenFolio': value.hasOpenFolio,
        'allowMobileCheckout': value.allowMobileCheckout,
        'allowMobileViewFolio': value.allowMobileViewFolio,
        'allowPreRegistration': value.allowPreRegistration,
        'optedForCommunication': value.optedForCommunication,
        'chargeCardNumber': value.chargeCardNumber,
    };
}
exports.HotelReservationTypeToJSON = HotelReservationTypeToJSON;
