/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Enum to denote the Status of Readiness messages sent to Guest Devices.
 * @export
 */
export declare const CommunicationStatusType: {
    readonly Pending: "Pending";
    readonly Completed: "Completed";
    readonly Failed: "Failed";
    readonly Sent: "Sent";
    readonly Received: "Received";
    readonly Cancelled: "Cancelled";
    readonly PendingAvailability: "PendingAvailability";
};
export type CommunicationStatusType = typeof CommunicationStatusType[keyof typeof CommunicationStatusType];
export declare function CommunicationStatusTypeFromJSON(json: any): CommunicationStatusType;
export declare function CommunicationStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommunicationStatusType;
export declare function CommunicationStatusTypeToJSON(value?: CommunicationStatusType | null): any;
